package display

import (
	"context"
	"fmt"

	"github.com/murashi19/koda-b8-ewallet-cli/internal/app"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/models"
)

func (m *UserMenu) Login(session *app.Session) {
	var req models.LoginRequest

	fmt.Println("===== Login =====")

	fmt.Print("Email    : ")
	fmt.Scan(&req.Email)

	fmt.Print("Password : ")
	fmt.Scan(&req.Password)

	user, err := m.userService.Login(
		context.Background(),
		req,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	session.Login(user)

	fmt.Printf("\nWelcome %s!\n", user.Name)
}
