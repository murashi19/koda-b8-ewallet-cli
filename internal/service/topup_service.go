package service

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/models"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/repo"
)

func (s *WalletService) TopUp(ctx context.Context, req models.TopUpRequest) error {
	if req.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	return s.withTransaction(
		ctx,
		func(walletRepo *repo.WalletRepository, transactionRepo *repo.TransactionRepository, tx pgx.Tx) error {

			wallet, err := walletRepo.GetWalletByUserID(
				ctx,
				req.UserID,
			)
			if err != nil {
				return err
			}

			transactionTypeID, err := transactionRepo.GetTransactionTypeIDByName(
				ctx,
				tx,
				"TOPUP",
			)
			if err != nil {
				return err
			}

			err = walletRepo.UpdateBalance(
				ctx,
				wallet.ID,
				req.Amount,
			)
			if err != nil {
				return err
			}

			transaction := models.Transaction{
				SenderWalletID:    nil,
				ReceiverWalletID:  &wallet.ID,
				TransactionTypeID: transactionTypeID,
				Amount:            req.Amount,
				Status:            "SUCCESS",
			}

			return transactionRepo.CreateTransaction(
				ctx,
				tx,
				transaction,
			)
		},
	)
}
