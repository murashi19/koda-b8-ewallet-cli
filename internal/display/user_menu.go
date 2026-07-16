package display

import (
	"context"
	"fmt"

	"github.com/murashi19/koda-b8-ewallet-cli/internal/models"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/service"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/utils"
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
	utils.ClearScreen()
}

func (m *UserMenu) ListUsers() {

	utils.ClearScreen()
	users, err := m.userService.GetAllUsers(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%-5s %-20s %-30s %-15s\n", "ID", "NAME", "EMAIL", "BALANCE")

	for _, user := range users {
		fmt.Printf(
			"%-5d %-20s %-30s Rp%d\n",
			user.ID,
			user.Name,
			user.Email,
			user.Balance,
		)
	}
	utils.EnterBack()
}
