package killbillApi

import (
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"
)

// Login Endpoint
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hit /Login endpoint. To be implemented")
}

// Signup Endpoint
func Signup(w http.ResponseWriter, r *http.Request) {
	// Endpoint for signing up
	var newUser User
	newUser.UUID = uuid.NewV4()

	fmt.Fprintln(w, "Hit /Signup endpoint. To be implemented")
	fmt.Printf("User uuid: %s\n", newUser.UUID)
}
