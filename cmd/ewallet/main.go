package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/murashi19/koda-b8-ewallet-cli/internal/app"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/config"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/display"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/service"
	"github.com/murashi19/koda-b8-ewallet-cli/internal/utils"
)

func main() {
	conn, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	done := make(chan struct{})
	go utils.Loading(done, "Database is being processed...")
	time.Sleep(2 * time.Second)
	close(done)

	fmt.Println("\nDatabase connected successfully ✅")
	time.Sleep(time.Second)

	session := &app.Session{}

	userService := service.NewUserService(conn)
	walletService := service.NewWalletService(conn)

	userMenu := display.NewUserMenu(userService)
	walletMenu := display.NewWalletMenu(walletService)

	for {
		utils.ClearScreen()

		if !session.IsLoggedIn() {
			printGuestMenu()
			choice := utils.ReadMenuChoice("Choose : ")

			switch choice {
			case 1:
				userMenu.CreateUser()
			case 2:
				userMenu.Login(session)
			case 0:
				fmt.Println("Thank you for using E-Wallet 👋")
				return
			default:
				fmt.Println("⚠️  Invalid menu, please enter a number from the list.")
				utils.EnterBack()
			}
			continue
		}

		printUserMenu(session)
		choice := utils.ReadMenuChoice("Choose : ")

		switch choice {
		case 1:
			walletMenu.ShowBalance(session)
		case 2:
			userMenu.ListUsers()
		case 3:
			walletMenu.TopUp(session)
		case 4:
			walletMenu.Withdraw(session)
		case 5:
			walletMenu.Transfer(session)
		case 6:
			walletMenu.TransactionHistory(session)
		case 7:
			session.Logout()
			fmt.Println("You have been logged out.")
			utils.EnterBack()
		default:
			fmt.Println("⚠️  Invalid menu, please enter a number from the list.")
			utils.EnterBack()
		}
	}
}

func printGuestMenu() {
	fmt.Println("========== KODA E-WALLET ==========")
	fmt.Println("1. Register")
	fmt.Println("2. Login")
	fmt.Println("0. Exit")
}

func printUserMenu(session *app.Session) {
	fmt.Println("==================================")
	fmt.Printf("Welcome, %s\n", session.CurrentUser.Name)
	fmt.Println("==================================")
	fmt.Println("1. Show Balance")
	fmt.Println("2. List Users")
	fmt.Println("3. Top Up")
	fmt.Println("4. Withdraw")
	fmt.Println("5. Transfer")
	fmt.Println("6. Transaction History")
	fmt.Println("7. Logout")
}
