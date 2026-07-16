package display

import (
	"context"
	"fmt"

	"github.com/murashi19/koda-b8-ewallet-cli/internal/app"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/models"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/utils"
)

func (m *WalletMenu) Withdraw(session *app.Session) {

	var req models.WithdrawRequest

	fmt.Println("===== Withdraw =====")

	req.UserID = session.CurrentUser.ID

	fmt.Print("Amount : ")
	fmt.Scan(&req.Amount)

	err := m.walletService.Withdraw(
		context.Background(),
		req,
	)
	if err != nil {
		fmt.Println()
		fmt.Println("Withdraw failed:", err)
		return
	}

	fmt.Println()
	fmt.Println("Withdraw success!")
	utils.EnterBack()

}
