package main

import (
	"testing"
)

func testCreateUser(t *testing.T) {
	// Test user creation works properly
	newUser := CreateUser("bobdylan@fake.com", "Bob", "Dylan")
	anotherUser := CreateUser("theman@fake.com", "The", "Man")
	if anotherUser.Email != "theman@fake.com" {
		t.Error("email is wrong")
	}
	if newUser.Firstname != "Bob" {
		t.Error("Firstname is wrong")
	}
	if newUser.Lastname != "Dylan" {
		t.Error("Lastname is wrong")
	}
	if newUser.UUID == anotherUser.UUID {
		t.Error("Different users have same UUID!")
	}
}
