package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	/* The following are the API endpoints. This may change over time*/

	router.HandleFunc("/login", Login)   // Endpoint for logging in
	router.HandleFunc("/signup", Signup) // Endpoint for signing up

	router.HandleFunc("/add-record/{accountId}", AddTransaction) // Endpoint for adding a transaction record

	router.HandleFunc("/get-record/{accountId}", GetTransactions)                 // Endpoint for getting all known transaction record in history
	router.HandleFunc("/get-record/{accountId}/{month}", GetTransactions)         // Endpoint for getting all transactions in a particular month
	router.HandleFunc("/get-record/{accountId}/{transactionId}", GetTransactions) //Endpoint for getting one particular transaction

	router.HandleFunc("/update-record/{accountId}/{transactionId}", UpdateTransaction)

	router.HandleFunc("/delete-record/{accountId}/{transactionId}", DeleteRecord) // Endpoint for deleting a transaction by its UID
	router.HandleFunc("/get-prediction/{accountId}", Predict)                     // Endpoint for retrieving a prediction for spending on a particular month

	log.Fatal(http.ListenAndServe(":8080", router))

}

// AddTransaction - This is the endpoint for adding a transaction to an account, which triggers updating balance of the given account
// the newly-created transaction's uuid is returned in response
func AddTransaction(w http.ResponseWriter, r *http.Request) {
	accountID := mux.Vars(r)["accountId"]
	fmt.Fprintln(w, "Hit /AddBalance endpoint. To be implemented", accountID)
}

// GetTransactions - This is the endpoint for getting one or more transactions associated with the given account
// response in array of transactions in json and status
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	accountID := mux.Vars(r)["accountId"]
	fmt.Fprintln(w, "Hit /GetTransactions endpoint. To be implemented", accountID)
}

// UpdateTransaction - This is the endpoint for updating a transaction, which triggers updating balance of the given account
// Looks at form data in *http.Request and use db services to update the specified transaction
func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	rvars := mux.Vars(r)
	accountID := rvars["accountId"]
	transactionID := rvars["transactionId"]
	fmt.Fprintln(w, "Hit /UpdateTransaction endpoint. To be implemented", accountID, transactionID)
}

// DeleteRecord - This is the endpoint for deleting a transaction, which also triggers updating balance of the given account
func DeleteRecord(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hit /DeleteRecord endpoint. To be implemented")
}

// Predict - Prediction endpoint
func Predict(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hit /Predict endpoint. To be implemented")
}
