package display

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/murashi19/koda-b8-ewallet-cli/internal/app"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/utils"
)

func (m *WalletMenu) TransactionHistory(session *app.Session) {

	utils.ClearScreen()

	userID := session.CurrentUser.ID

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
