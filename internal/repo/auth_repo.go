package repo

import (
	"context"
	"fmt"

	"github.com/murashi19/koda-b8-ewallet-cli/internal/models"
)

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	query := `
		SELECT
			id,
			name,
			email,
			password,
			created_at,
			updated_at
		FROM users
		WHERE email = $1;
	`

	var user models.User

	err := r.db.QueryRow(
		ctx,
		query,
		email,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return models.User{}, fmt.Errorf("get user by email: %w", err)
	}

	return user, nil
}
