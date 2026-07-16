package display

import (
	"context"
	"fmt"

	"github.com/murashi19/koda-b8-ewallet-cli/internal/models"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/utils"
)

func (m *WalletMenu) TopUp() {

	utils.ClearScreen()

	var req models.TopUpRequest

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
