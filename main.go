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

	router.HandleFunc("/add-record/{accountId}", AddRecord) // Endpoint for adding a transaction record

	router.HandleFunc("/get-record/{accountId}/{month}", AddRecord)     // Endpoint for getting all known transaction record in history
	router.HandleFunc("/get-record/{accountId}/{month}", AddRecord)     // Endpoint for adding transaction record in particular month
	router.HandleFunc("/get-record/{accountId}/{from}/{to}", AddRecord) // Endpoint for adding transaction record from a particular month to another

	router.HandleFunc("/delete-record/{transactionId}", DeleteRecord) // Endpoint for deleting a transaction by its UID
	router.HandleFunc("/get-prediction/{accountId}", Predict)         // Endpoint for retrieving a prediction for spending on a particular month

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hit /Login endpoint. To be implemented")
}

func Signup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hit /Signup endpoint. To be implemented")
}

func AddRecord(w http.ResponseWriter, r *http.Request) {
	accountId := mux.Vars(r)["accountId"]
	fmt.Fprintln(w, "Hit /AddRecord endpoint. To be implemented", accountId)
}

func DeleteRecord(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hit /DeleteRecord endpoint. To be implemented")
}

func Predict(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hit /Predict endpoint. To be implemented")
}
