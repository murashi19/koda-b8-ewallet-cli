package service

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/models"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/repo"
)

type WalletService struct {
	db *pgx.Conn
}

func NewWalletService(db *pgx.Conn) *WalletService {
	return &WalletService{
		db: db,
	}
}

func (s *WalletService) repositories(tx pgx.Tx) (
	*repo.WalletRepository,
	*repo.TransactionRepository,
) {
	return repo.NewWalletRepository(tx),
		repo.NewTransactionRepository(tx)
}

func (s *WalletService) GetWalletByUserID(ctx context.Context, userID int64) (models.Wallet, error) {
	walletRepo := repo.NewWalletRepository(s.db)
	wallet, err := walletRepo.GetWalletByUserID(ctx, userID)
	if err != nil {
		return models.Wallet{}, err
	}
	return wallet, nil
}

func (s *WalletService) GetWalletDetailByUserID(ctx context.Context, userID int64) (models.WalletDetail, error) {
	walletRepo := repo.NewWalletRepository(s.db)
	return walletRepo.GetWalletDetailByUserID(ctx, userID)
}

func (s *WalletService) withTransaction(ctx context.Context,
	fn func(walletRepo *repo.WalletRepository, transactionRepo *repo.TransactionRepository, tx pgx.Tx,
	) error) error {

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	walletRepo := repo.NewWalletRepository(tx)
	transactionRepo := repo.NewTransactionRepository(tx)

	if err := fn(
		walletRepo,
		transactionRepo,
		tx,
	); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (s *WalletService) recordTransaction(
	ctx context.Context,
	tx pgx.Tx,
	transactionRepo *repo.TransactionRepository,
	transactionType string,
	sender *int64,
	receiver *int64,
	amount int64,
) error {

	transactionTypeID, err := transactionRepo.GetTransactionTypeIDByName(
		ctx,
		tx,
		transactionType,
	)
	if err != nil {
		return err
	}

	transaction := models.Transaction{
		SenderWalletID:    sender,
		ReceiverWalletID:  receiver,
		TransactionTypeID: transactionTypeID,
		Amount:            amount,
		Status:            "SUCCESS",
	}

	return transactionRepo.CreateTransaction(
		ctx,
		tx,
		transaction,
	)
}
