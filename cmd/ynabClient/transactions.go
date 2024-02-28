package ynabClient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type TransactionService struct{}

type transaction struct {
	Id                      string  `json:"id"`
	Date                    string  `json:"date"`
	Amount                  int     `json:"amount"`
	Memo                    *string `json:"memo"`
	Cleared                 string  `json:"cleared"`
	Approved                bool    `json:"approved"`
	FlagColor               *string `json:"flag_color"`
	FlagName                *string `json:"flag_name"`
	AccountID               string  `json:"account_id"`
	PayeeID                 *string `json:"payee_id"`
	CategoryID              *string `json:"category_id"`
	TransferAccountID       *string `json:"transfer_account_id"`
	TransferTransactionID   *string `json:"transfer_transaction_id"`
	MatchedTransactionID    *string `json:"matched_transaction_id"`
	ImportID                *string `json:"import_id"`
	ImportPayeeName         *string `json:"import_payee_name"`
	ImportPayeeNameOriginal *string `json:"import_payee_name_original"`
	DebtTransactionType     *string `json:"debt_transaction_type"`
	Deleted                 bool    `json:"deleted"`
	AccountName             string  `json:"account_name"`
	PayeeName               *string `json:"payee_name"`
	CategoryName            *string `json:"category_name"`
	// TODO: add Subtransactions
}

type transactionListResponse struct {
	Data struct {
		Transactions []transaction `json:"transactions"`
		// TODO: add ServerKnowledge
	} `json:"data"`
}

func (t *TransactionService) List() (response *transactionListResponse, err error) {
	url := fmt.Sprintf("%s/budgets/last-used/transactions", YNAB_API_URL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+os.Getenv("YNAB_API_KEY"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
