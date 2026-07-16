package main

import (
	"context"
	"fmt"
	"log"
	"time"

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
	fmt.Println("")

	close(done)
	fmt.Println("Database connected successfully ✅")
	time.Sleep(1 * time.Second)

	userService := service.NewUserService(conn)

	menu := display.NewUserMenu(userService)

	for {
		utils.ClearScreen()
		fmt.Println("===== E-Wallet =====")
		fmt.Println("1. Create User")
		fmt.Println("0. Exit")

		var choose int

		fmt.Print("Choose : ")
		fmt.Scan(&choose)

		switch choose {

		case 1:
			menu.CreateUser()

		case 0:
			fmt.Println("Thank you for using E-Wallet Apps Bye👋")
			return

		default:
			fmt.Println("Invalid menu")
		}

	}
}
