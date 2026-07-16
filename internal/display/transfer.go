package display

import (
	"context"
	"fmt"

	"github.com/murashi19/koda-b8-ewallet-cli/internal/app"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/models"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/utils"
)

func (m *WalletMenu) Transfer(session *app.Session) {

	var req models.TransferRequest

	fmt.Println("===== Transfer =====")

	req.SenderUserID = session.CurrentUser.ID

	fmt.Print("Receiver User ID : ")
	fmt.Scan(&req.ReceiverUserID)

	fmt.Print("Amount           : ")
	fmt.Scan(&req.Amount)

	err := m.walletService.Transfer(
		context.Background(),
		req,
	)
	if err != nil {
		fmt.Println()
		fmt.Println("Transfer failed:", err)
		return
	}

	fmt.Println()
	fmt.Println("Transfer success!")
	utils.EnterBack()
}
