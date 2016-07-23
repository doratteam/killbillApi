package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/satori/go.uuid"
)

type User struct {
	UUID      uuid.UUID
	Email     string
	Firstname string
	Lastname  string
}

func CreateUser(email string, firstname string, lastname string) User {
	// Make a new User instance
	return User{
		uuid.NewV4(),
		email,
		firstname,
		lastname,
	}
}
