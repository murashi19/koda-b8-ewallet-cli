package display

import (
	"context"
	"fmt"

	"github.com/murashi19/koda-b8-ewallet-cli/internal/app"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/models"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/utils"
)

func (m *WalletMenu) TopUp(session *app.Session) {

	utils.ClearScreen()

	var req models.TopUpRequest

	req.UserID = session.CurrentUser.ID
	fmt.Println("===== TOP UP =====")

	fmt.Print("User ID : ")
	fmt.Scan(&req.UserID)

	fmt.Print("Amount : ")
	fmt.Scan(&req.Amount)

	err := m.walletService.TopUp(
		context.Background(),
		req,
	)

	if err != nil {
		fmt.Println("Error:", err)
		utils.EnterBack()
		return
	}

	fmt.Println()
	fmt.Println("Top Up Success!")

	utils.EnterBack()
}
