package service

import (
	"context"
	"errors"

	"github.com/murashi19/koda-b8-ewallet-cli/internal/models"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/repo"
)

func (s *WalletService) Transfer(ctx context.Context, req models.TransferRequest) error {

	// Validation
	if req.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	if req.SenderUserID == req.ReceiverUserID {
		return errors.New("cannot transfer to yourself")
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

	// Sender Wallet
	senderWallet, err := walletRepo.GetWalletByUserID(
		ctx,
		req.SenderUserID,
	)
	if err != nil {
		return err
	}

	// Receiver Wallet
	receiverWallet, err := walletRepo.GetWalletByUserID(
		ctx,
		req.ReceiverUserID,
	)
	if err != nil {
		return err
	}

	// Balance Validation
	if senderWallet.Balance < req.Amount {
		return errors.New("insufficient balance")
	}

	// Transaction Type
	transactionTypeID, err := transactionRepo.GetTransactionTypeIDByName(
		ctx,
		tx,
		"TRANSFER",
	)
	if err != nil {
		return err
	}

	// Deduct Sender
	err = walletRepo.UpdateBalance(
		ctx,
		senderWallet.ID,
		-req.Amount,
	)
	if err != nil {
		return err
	}

	// Add Receiver
	err = walletRepo.UpdateBalance(
		ctx,
		receiverWallet.ID,
		req.Amount,
	)
	if err != nil {
		return err
	}

	// Save Transaction
	transaction := models.Transaction{
		SenderWalletID:    &senderWallet.ID,
		ReceiverWalletID:  &receiverWallet.ID,
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
