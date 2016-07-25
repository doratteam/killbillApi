package main

import (
	"database/sql"
	"strconv"

	"github.com/satori/go.uuid"
)

// Transaction - model for transactions
type Transaction struct {
	UUID       uuid.UUID
	AccountID  uuid.UUID
	Date       string // DATE column in SQL?
	Amount     float64
	Categories int64 //TODO: need to figure out how we're going to represent categories for Transaction
	Desc       string
}

// Columns - called by db service to get list of columns for this model
func (t Transaction) Columns() string {
	//TODO: these are temporary column names. change this function to reflect actual column names
	return "(`uuid`,`account_id`,`t_date`,`amount`,`categories`,`desc`)"
}

// Serialize - called by db service to get values for each column when saving/updating
func (t Transaction) Serialize() string {
	//TODO: remember to modify this when changing model/db table
	return "(" + t.UUID.String() + "," + t.AccountID.String() + "," + t.Date + "," + strconv.FormatFloat(t.Amount, 'f', 2, 64) + "," + strconv.FormatInt(t.Categories, 10) + "," + t.Desc + ")"
}

// DeSerialize - called by db service to deserialize a row to a Transaction
func (t Transaction) DeSerialize(r *sql.Row) interface{} {
	//TODO: scan the given row with r.Scan() function and return an appropriate Transaction
	return Transaction{}
}

// CreateTransaction - creates a new transaction and returns it
func CreateTransaction(accountID uuid.UUID, date string, amount float64, categories int64, desc string) Transaction {
	return Transaction{
		uuid.NewV4(),
		accountID,
		date,
		amount,
		categories,
		desc,
	}
}

// TODO: Below are functions that need to be implemented in model for our endpoints to fully function

// SaveTransaction - saves a transaction to db
func (t Transaction) SaveTransaction(db *sql.DB) {
	//This will call db service to save a transaction
}

// UpdateTransaction - updates a transaction
func (t Transaction) UpdateTransaction(db *sql.DB) {

}

func SelectTransactionWithID(transactionID string) Transaction {
	return Transaction{}
}

func SelectTransactionsForAccount(accountID string) []Transaction {
	return []Transaction{}
}

func SelectTransactionsForAccountWithDateRange(accountID string, start string, end string) []Transaction {
	return []Transaction{}
}
