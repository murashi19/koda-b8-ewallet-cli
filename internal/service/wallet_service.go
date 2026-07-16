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

	// transaction
	transactionTypeID, err := transactionRepo.GetTransactionTypeIDByName(
		ctx,
		tx,
		"TOPUP",
	)
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

	// Save transaction

	transaction := models.Transaction{
		SenderWalletID:    nil,
		ReceiverWalletID:  &wallet.ID,
		TransactionTypeID: transactionTypeID,
		Amount:            req.Amount,
		Status:            "SUCCESS",
	}

	err = transactionRepo.CreateTransaction(ctx, tx, transaction)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *WalletService) GetTransactionHistory(
	ctx context.Context,
	userID int64,
) ([]models.TransactionHistory, error) {

	walletRepo := repo.NewWalletRepository(s.db)
	transactionRepo := repo.NewTransactionRepository(s.db)

	wallet, err := walletRepo.GetWalletByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	details, err := transactionRepo.GetTransactionByWalletID(
		ctx,
		wallet.ID,
	)
	if err != nil {
		return nil, err
	}

	var histories []models.TransactionHistory

	for _, detail := range details {

		history := models.TransactionHistory{
			Type:      detail.Type,
			Amount:    detail.Amount,
			Status:    detail.Status,
			CreatedAt: detail.CreatedAt,
		}

		// Business Logic
		if detail.Type == "TRANSFER" {

			if detail.SenderWalletID != nil &&
				*detail.SenderWalletID == wallet.ID {

				history.Type = "TRANSFER OUT"

			} else {

				history.Type = "TRANSFER IN"
			}
		}

		histories = append(histories, history)
	}

	return histories, nil
}

func (s *WalletService) Withdraw(ctx context.Context, req models.WithdrawRequest) error {

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

	// Business Validation
	if wallet.Balance < req.Amount {
		return errors.New("insufficient balance")
	}

	// Transaction Type
	transactionTypeID, err := transactionRepo.GetTransactionTypeIDByName(
		ctx,
		tx,
		"WITHDRAW",
	)
	if err != nil {
		return err
	}

	// Update Balance (kurangi saldo)
	err = walletRepo.UpdateBalance(
		ctx,
		wallet.ID,
		-req.Amount,
	)
	if err != nil {
		return err
	}

	// Save Transaction
	transaction := models.Transaction{
		SenderWalletID:    &wallet.ID,
		ReceiverWalletID:  nil,
		TransactionTypeID: transactionTypeID,
		Amount:            req.Amount,
		Status:            "SUCCESS",
	}

	err = transactionRepo.CreateTransaction(
		ctx,
		tx,
		transaction,
	)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
