package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const YNAB_API_URL = "https://api.ynab.com/v1/"

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello, World!")
}

func budgetHandler(w http.ResponseWriter, r *http.Request) {
	http.Header.Set("Authorization", "Bearer "+os.Getenv("YNAB_API_KEY"))
	budgets, err := http.Get(YNAB_API_URL + "budgets/")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, "Budgets", budgets)
}

func transactionsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Transactions")
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/transactions", transactionsHandler)
	mux.HandleFunc("/budgets", budgetHandler)

	fmt.Println("Server is running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
