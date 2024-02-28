package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ghernandez345/ynab-hub/cmd/ynabClient"
	"github.com/joho/godotenv"
)

func prtSupportedServices() {
	fmt.Println("Supported services:")
	fmt.Println("  - transactions")
}

func main() {

	// auth token setup
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	token := os.Getenv("YNAB_API_KEY")

	client, err := ynabClient.NewClient(token)
	if err != nil {
		fmt.Println("Error creating client: ", err)
		return
	}

	service := os.Args[1]

	switch service {
	case "transactions":
		fmt.Println("getting transactions...")
		transactions, err := client.Transactions.List()
		if err != nil {
			fmt.Println("Error getting transactions: ", err)
			return
		}
		fmt.Println(transactions)
	default:
		fmt.Println("This service is not currently supported.")
		prtSupportedServices()
	}
}
