package repo

import (
	"context"
	"errors"
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

	_, err := r.db.Exec(ctx,
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

func (r *WalletRepository) GetWalletByUserID(ctx context.Context, userID int64) (models.Wallet, error) {
	query := `SELECT id, user_id, balance, currency, created_at, updated_at
		FROM wallets
		WHERE user_id = $1;`

	var wallet models.Wallet

	err := r.db.QueryRow(ctx,
		query,
		userID,
	).Scan(
		&wallet.ID,
		&wallet.UserID,
		&wallet.Balance,
		&wallet.Currency,
		&wallet.CreatedAt,
		&wallet.UpdatedAt,
	)

	if err != nil {
		return models.Wallet{}, fmt.Errorf("get wallet by user id: %w", err)
	}

	return wallet, nil
}

func (r *WalletRepository) UpdateBalance(ctx context.Context, walletID int64, delta int64) error {
	query := `
		UPDATE wallets
		SET
			balance = balance + $1,
			updated_at = CURRENT_TIMESTAMP
		WHERE
			id = $2
			AND balance + $1 >= 0;`

	result, err := r.db.Exec(ctx,
		query,
		delta,
		walletID,
	)
	if err != nil {
		return fmt.Errorf("update wallet balance: %w", err)
	}

	if result.RowsAffected() == 0 {
		return errors.New("wallet not found or insufficient balance")
	}

	return nil
}

func (r *WalletRepository) GetWalletDetailByUserID(ctx context.Context, userID int64) (models.WalletDetail, error) {
	query := `
		SELECT u.id, u.name, w.balance, w.currency
		FROM wallets AS w
		INNER JOIN users AS u
			ON w.user_id = u.id
		WHERE u.id = $1;`

	var wallet models.WalletDetail

	err := r.db.QueryRow(ctx,
		query,
		userID,
	).Scan(
		&wallet.UserID,
		&wallet.UserName,
		&wallet.Balance,
		&wallet.Currency,
	)

	if err != nil {
		return models.WalletDetail{}, fmt.Errorf("get wallet detail: %w", err)
	}

	return wallet, nil
}
