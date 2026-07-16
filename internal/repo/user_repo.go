package repo

import (
	"context"

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
