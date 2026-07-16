package display

import (
	"fmt"

	"github.com/murashi19/koda-b8-ewallet-cli/internal/app"
)

func Logout(session *app.Session) {
	session.Logout()
	fmt.Println("Logout success.")
}
