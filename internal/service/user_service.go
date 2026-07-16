package service

import (
	"context"
	"errors"
	"fmt"
	"net/mail"

	"github.com/jackc/pgx/v5"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/models"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/repo"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	db *pgx.Conn
}

func NewUserService(db *pgx.Conn) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]models.UserBalance, error) {
	repo := repo.NewUserRepository(s.db)

	users, err := repo.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) CreateUser(ctx context.Context, req models.RegisterRequest) error {

	// Validation
	if req.Name == "" {
		fmt.Println("isian")
		utils.EnterBack()
		return errors.New("name is required")
	}

	if req.Email == "" {
		fmt.Println("isian emailna")
		utils.EnterBack()
		return errors.New("email is required")
	}

	if _, err := mail.ParseAddress(req.Email); err != nil {
		fmt.Println("email invalid")
		utils.EnterBack()
		return errors.New("invalid email format")
	}

	if req.Password == "" {
		fmt.Println("password salah")
		utils.EnterBack()
		return errors.New("password is required")
	}

	if len(req.Password) < 6 {
		fmt.Println("password kurang")
		utils.EnterBack()
		return errors.New("password must be at least 6 characters")
	}

	if req.PhoneNumber == "" {
		fmt.Println("no hp invalid")
		utils.EnterBack()
		return errors.New("phone number is required")
	}

	tx, err := s.db.Begin(ctx)
	if err != nil {
		fmt.Println("nul.....")
		return err
	}
	defer tx.Rollback(ctx)

	userRepo := repo.NewUserRepository(tx)
	walletRepo := repo.NewWalletRepository(tx)

	// Check Email
	exists, err := userRepo.ExistsByEmail(ctx, req.Email)
	if err != nil {
		fmt.Println("cek email?")
		return err
	}

	if exists {
		return errors.New("email already registered")
	}

	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		fmt.Println("HIdden pass")
		utils.EnterBack()
		return err
	}

	// Entity
	user := models.User{
		Name:        req.Name,
		Email:       req.Email,
		Password:    string(hashedPassword),
		PhoneNumber: req.PhoneNumber,
	}

	userID, err := userRepo.CreateUser(ctx, user)
	if err != nil {
		fmt.Println("halah error")
		utils.EnterBack()
		return err
	}

	wallet := models.Wallet{
		UserID:   userID,
		Balance:  models.DefaultBalance,
		Currency: models.DefaultCurrency,
	}

	if err := walletRepo.CreateWallet(ctx, wallet); err != nil {
		fmt.Println("ssdasdasd")
		utils.EnterBack()
		return err
	}

	return tx.Commit(ctx)
}
