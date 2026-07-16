package main

import (
	"context"
	"fmt"
	"log"

	"github.com/murashi19/koda-b8-ewallet-cli/internal/config"
)

func main() {
	conn, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	fmt.Println("Database connected successfully 🚀")
}
