package service

import (
	"context"
	"errors"

	"github.com/murashi19/koda-b8-ewallet-cli/internal/models"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/repo"
)

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
