package display

import (
	"context"
	"fmt"

	"github.com/murashi19/koda-b8-ewallet-cli/internal/service"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/utils"
)

type WalletMenu struct {
	walletService *service.WalletService
}

func NewWalletMenu(walletService *service.WalletService) *WalletMenu {
	return &WalletMenu{
		walletService: walletService,
	}
}
func (m *WalletMenu) ShowBalance() {

	utils.ClearScreen()

	var userID int64

	fmt.Println("===== SHOW BALANCE =====")
	fmt.Print("User ID : ")
	fmt.Scan(&userID)

	wallet, err := m.walletService.GetWalletByUserID(
		context.Background(),
		userID,
	)

	if err != nil {
		fmt.Println("Error:", err)
		utils.EnterBack()
		return
	}

	fmt.Println()
	fmt.Println("===== WALLET =====")
	fmt.Printf("User ID  : %s\n", wallet.UserName)
	fmt.Printf("Balance  : Rp%d\n", wallet.Balance)
	fmt.Printf("Currency : %s\n", wallet.Currency)

	utils.EnterBack()
}
