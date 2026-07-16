package service

import (
	"context"
	"errors"

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

func (s *WalletService) TopUp(ctx context.Context, req models.TopUpRequest) error {

	// Validation
	if req.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	// Begin Transaction
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	// Repository
	walletRepo := repo.NewWalletRepository(tx)
	transactionRepo := repo.NewTransactionRepository(tx)

	// Get Wallet
	wallet, err := walletRepo.GetWalletByUserID(ctx, req.UserID)
	if err != nil {
		return err
	}
	// Update Balance
	err = walletRepo.UpdateBalance(
		ctx,
		wallet.ID,
		req.Amount,
	)
	if err != nil {
		return err
	}

	// TODO:
	// Save transaction

	_ = transactionRepo

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
