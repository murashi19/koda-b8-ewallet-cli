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
