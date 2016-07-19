package view

import (
	"fmt"
	"net/http"

	"model"

	"github.com/satori/go.uuid"
)

// Login Endpoint
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hit /Login endpoint. To be implemented")
}

// Signup Endpoint
func Signup(w http.ResponseWriter, r *http.Request) {
	// Endpoint for signing up
	var newUser model.User
	newUser.UUID = uuid.NewV4()

	fmt.Fprintln(w, "Hit /Signup endpoint. To be implemented")
	fmt.Printf("User uuid: %s\n", newUser.UUID)
}
