package main

type ClearedStatus int

const (
	Cleared ClearedStatus = iota
	Uncleared
	Reconciled
)

type Transaction struct {
	Id           string
	Date         string
	Amount       int64
	Memo         *string
	Cleared      ClearedStatus // TODO: make sure this works
	Approved     bool
	FlatColor    string
	AccountId    string
	PayeeName    string
	CategoryName string
}
