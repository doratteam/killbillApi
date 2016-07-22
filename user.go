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

func CommitUser(user User) {
	fmt.Println("Not implemented")
}

func getUserDB() *sql.DB {
	// Returns pointer to the sql.DB object that can be used
	// 07/22/16 jameswhang
	// NOTE #1: sql.DB instance is long-living. Don't close this shit until the program goes down
	// NOTE #2: this isn't actually connecting to the database, but rather a 'prep' for the driver to open it up.
	// Read http://go-database-sql.org/accessing.html for more details

	db, err := sql.Open("mysql", "kb:kb@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create User table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS KbUsers ( uuid integer PRIMARY KEY, email varchar(255), firstname varchar(32), lastname varchar(32)")
	if err != nil {
		panic(err)
	}
	return db
}

func closeUserDB(db sql.DB) {
	db.Close()
}
