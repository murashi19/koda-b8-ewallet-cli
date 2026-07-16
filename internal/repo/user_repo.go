package repo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/models"
)

type UserRepository struct {
	db DBTX
}

func NewUserRepository(db DBTX) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// GET ALL USER
func (r *UserRepository) GetAllUsers(ctx context.Context) ([]models.UserBalance, error) {
	query := `
	SELECT
		u.id,
		u.name,
		u.email,
		w.balance
	FROM users u
	INNER JOIN wallets w
		ON u.id = w.user_id
	ORDER BY u.id;
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query users: %w", err)
	}
	defer rows.Close()

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.UserBalance])
	if err != nil {
		return nil, fmt.Errorf("collect users: %w", err)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return users, nil
}

// CREATE USER
func (r *UserRepository) CreateUser(ctx context.Context, user models.User) (int64, error) {
	var id int64
	query := `INSERT INTO users ( name, email, password, phone_number )
		VALUES ($1, $2, $3, $4)
		RETURNING id;`

	err := r.db.QueryRow(context.Background(),
		query,
		user.Name,
		user.Email,
		user.Password,
		user.PhoneNumber,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
