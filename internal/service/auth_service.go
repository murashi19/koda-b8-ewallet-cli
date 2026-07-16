package service

import (
	"context"
	"errors"

	"github.com/murashi19/koda-b8-ewallet-cli/internal/models"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/repo"
	"golang.org/x/crypto/bcrypt"
)

func (s *UserService) Login(ctx context.Context, req models.LoginRequest) (models.User, error) {
	userRepo := repo.NewUserRepository(s.db)
	user, err := userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return models.User{}, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)

	if err != nil {
		return models.User{}, errors.New("invalid email or password")
	}

	return user, nil
}
