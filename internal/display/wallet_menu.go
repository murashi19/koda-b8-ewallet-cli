package display

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/murashi19/koda-b8-ewallet-cli/internal/models"
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

	wallet, err := m.walletService.GetWalletDetailByUserID(
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
	fmt.Printf("Balance  : Rp %d\n", wallet.Balance)
	fmt.Printf("Currency : %s\n", wallet.Currency)

	utils.EnterBack()
}

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

func (m *WalletMenu) TransactionHistory() {

	utils.ClearScreen()
	var userID int64

	fmt.Print("User ID : ")
	fmt.Scanln(&userID)

	histories, err := m.walletService.GetTransactionHistory(
		context.Background(),
		userID,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	if len(histories) == 0 {
		fmt.Println("No transaction history found.")
		return
	}

	fmt.Println("\n=== Transaction History ===")

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "No\tDate\tType\tAmount")
	fmt.Fprintln(w, "--\t----\t----\t------")

	for i, h := range histories {
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\n",
			i+1,
			h.CreatedAt.Format("02-01-2006 15:04"),
			h.Type,
			formatRupiah(h.Amount),
		)
	}

	w.Flush()
	fmt.Println("")
	utils.EnterBack()
}

func formatRupiah(amount int64) string {
	return fmt.Sprintf("Rp%d", amount)
}

func (m *WalletMenu) Withdraw() {

	var req models.WithdrawRequest

	fmt.Println("===== Withdraw =====")

	fmt.Print("User ID : ")
	fmt.Scan(&req.UserID)

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
