package ynabClient

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Client struct {
	Transactions Service[Transaction]
}

const YNAB_API_URL = "https://api.ynab.com/v1/"

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello, World!")
}

func budgetHandler(w http.ResponseWriter, r *http.Request) {
	url := YNAB_API_URL + "budgets/"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+os.Getenv("YNAB_API_KEY"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, string(body))
}

func transactionDetailsHandler(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/budgets/%s/transactions/%s", YNAB_API_URL, "a434d13b-be02-4d60-9e4d-0fa45ea8c744", r.PathValue("id"))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+os.Getenv("YNAB_API_KEY"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, string(body))
}

func NewClient(token string) (client *Client, err error) {

	client = &Client{
		Transactions: &TransactionService{},
	}

	return client, nil
}
