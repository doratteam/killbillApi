package main

import (
	"fmt"
	"net/http"
)

// Login Endpoint
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hit /Login endpoint. To be implemented")
}

// Signup Endpoint
func Signup(w http.ResponseWriter, r *http.Request) {
	// Endpoint for signing up

	fmt.Fprintln(w, "Hit /Signup endpoint. To be implemented")
}
