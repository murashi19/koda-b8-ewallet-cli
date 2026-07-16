package service

import (
	"context"

	"github.com/murashi19/koda-b8-ewallet-cli/internal/models"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/repo"
)

func (s *WalletService) GetTransactionHistory(ctx context.Context, userID int64) ([]models.TransactionHistory, error) {

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
