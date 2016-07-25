package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/satori/go.uuid"
)

// Serializer - interface every model should implement for db service to be able to save them to db
type Serializer interface {
	Serialize() string
}

// ListColumn - interface every model should implement for db service to know which columns exist for the model's table
type ListColumn interface {
	Columns() string
}

type User struct {
	UUID      uuid.UUID
	Email     string
	Firstname string
	Lastname  string
}

func OpenDatabases() *sql.DB {
	// Returns pointer to the sql.DB object that can be used
	// 07/22/16 jameswhang
	// NOTE #1: sql.DB instance is long-living. Don't close this shit until the program goes down
	// NOTE #2: this isn't actually connecting to the database, but rather a 'prep' for the driver to open it up.
	// Read http://go-database-sql.org/accessing.html for more details

	db, err := sql.Open("mysql", "kb:kb@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}

	// Create User table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS KbUsers ( uuid integer NOT NULL AUTO_INCREMENT, email varchar(255), firstname varchar(32), lastname varchar(32), PRIMARY_KEY (uuid)")
	if err != nil {
		panic(err)
	}

	// Create Transactions table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS KbTransactions ( uuid integer NOT NULL AUTO_INCREMENT, ownerUUID integer, amount float(9) NOT NULL, PRIMARY_KEY(uuid), FOREIGN KEY (ownerUUID) REFERENCES kbUsers(uuid)")
	if err != nil {
		panic(err)
	}
	return db
}

func closeUserDB(db sql.DB) {
	db.Close()
}
