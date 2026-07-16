package display

import (
	"context"
	"fmt"

	"github.com/murashi19/koda-b8-ewallet-cli/internal/models"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/service"
)

type UserMenu struct {
	userService *service.UserService
}

func NewUserMenu(userService *service.UserService) *UserMenu {
	return &UserMenu{
		userService: userService,
	}
}

func (m *UserMenu) CreateUser() {

	var user models.User

	fmt.Print("Name : ")
	fmt.Scan(&user.Name)

	fmt.Print("Email : ")
	fmt.Scan(&user.Email)

	fmt.Print("Password : ")
	fmt.Scan(&user.Password)

	fmt.Print("Phone Number : ")
	fmt.Scan(&user.PhoneNumber)

	err := m.userService.CreateUser(context.Background(), user)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	fmt.Println("User created successfully.")
}
