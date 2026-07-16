package repo

import (
	"context"
	"fmt"

	"github.com/murashi19/koda-b8-ewallet-cli/internal/models"
)

type WalletRepository struct {
	db DBTX
}

func NewWalletRepository(db DBTX) *WalletRepository {
	return &WalletRepository{
		db: db,
	}
}

func (r *WalletRepository) CreateWallet(ctx context.Context, wallet models.Wallet) error {

	query := `INSERT INTO wallets (user_id,balance,currency)
		VALUES ($1,$2,$3);`

	_, err := r.db.Exec(
		ctx,
		query,
		wallet.UserID,
		wallet.Balance,
		wallet.Currency,
	)

	if err != nil {
		return fmt.Errorf("create wallet: %w", err)
	}

	return nil
}

func (r *WalletRepository) GetWalletByUserID(ctx context.Context, userID int64) (models.WalletDetail, error) {
	query := `
		SELECT
			u.name,
			w.balance,
			w.currency
		FROM wallets AS w
		INNER JOIN users AS u
			ON u.id = w.user_id
		WHERE w.user_id = $1;
	`

	var wallet models.WalletDetail

	err := r.db.QueryRow(
		ctx,
		query,
		userID,
	).Scan(
		&wallet.UserName,
		&wallet.Balance,
		&wallet.Currency,
	)

	if err != nil {
		return models.WalletDetail{}, fmt.Errorf("get wallet by user id: %w", err)
	}

	return wallet, nil
}
