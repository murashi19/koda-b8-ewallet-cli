package repo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/models"
)

type TransactionRepository struct {
	db DBTX
}

func NewTransactionRepository(db DBTX) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}
func (r *TransactionRepository) CreateTransaction(ctx context.Context, tx pgx.Tx, transaction models.Transaction) error {

	_, err := tx.Exec(
		ctx,
		`
		INSERT INTO transactions (
			sender_wallet_id,
			receiver_wallet_id,
			transaction_type_id,
			amount,
			status
		)
		VALUES ($1, $2, $3, $4, $5)
		`,
		transaction.SenderWalletID,
		transaction.ReceiverWalletID,
		transaction.TransactionTypeID,
		transaction.Amount,
		transaction.Status,
	)

	return err
}

func (r *TransactionRepository) GetTransactionTypeIDByName(ctx context.Context, tx pgx.Tx, name string) (int64, error) {

	var id int64
	err := tx.QueryRow(
		ctx,
		`
        SELECT id
        FROM transaction_types
        WHERE name = $1
        `,
		name,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *TransactionRepository) GetTransactionByWalletID(
	ctx context.Context,
	walletID int64,
) ([]models.TransactionDetail, error) {

	rows, err := r.db.Query(
		ctx,
		`
        SELECT
            tt.name,
            t.sender_wallet_id,
            t.receiver_wallet_id,
            t.amount,
            t.status,
            t.created_at
        FROM transactions t
        JOIN transaction_types tt
            ON tt.id = t.transaction_type_id
        WHERE
            t.sender_wallet_id = $1
            OR
            t.receiver_wallet_id = $1
        ORDER BY t.created_at DESC
        `,
		walletID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var histories []models.TransactionDetail

	for rows.Next() {

		var history models.TransactionDetail

		err := rows.Scan(
			&history.Type,
			&history.SenderWalletID,
			&history.ReceiverWalletID,
			&history.Amount,
			&history.Status,
			&history.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		histories = append(histories, history)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return histories, nil
}
