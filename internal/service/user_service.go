package service

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/models"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/repo"
)

type UserService struct {
	db *pgx.Conn
}

func NewUserService(
	db *pgx.Conn,
) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user models.User) error {

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	userRepo := repo.NewUserRepository(tx)
	walletRepo := repo.NewWalletRepository(tx)

	userID, err := userRepo.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	wallet := models.Wallet{
		UserID:   userID,
		Balance:  0,
		Currency: "IDR",
	}

	err = walletRepo.CreateWallet(ctx, wallet)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
