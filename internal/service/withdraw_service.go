package service

import (
	"context"
	"errors"

	"github.com/murashi19/koda-b8-ewallet-cli/internal/models"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/repo"
)

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
