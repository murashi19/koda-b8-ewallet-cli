package display

import (
	"context"
	"fmt"

	"github.com/murashi19/koda-b8-ewallet-cli/internal/models"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/utils"
)

func (m *WalletMenu) Transfer() {

	var req models.TransferRequest

	fmt.Println("===== Transfer =====")

	fmt.Print("Sender User ID   : ")
	fmt.Scan(&req.SenderUserID)

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
