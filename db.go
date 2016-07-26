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
	DeSerialize(r *sql.Row) interface{}
}

// ListColumn - interface every model should implement for db service to know which columns exist for the model's table
// on a second thought, this might be unnecessary...
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

func insertToDB(db sql.DB, tableName string, dbQuery string) int64 {
	// Takes in pointer to the db, table name as String, and the query to execute as string
	// Returns the ID of the row
	stmt, err := db.Prepare("INSERT INTO " + tableName + " VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(dbQuery)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return lastId
}

func readFromDB(db sql.DB, tablename string, fields string, query string, value string) string {
	// Reads from User DB that matches the query and returns string of result
	// NOTE: go/sql Scan function already implements the deserialization necessary
	// Just means DB layer needs to contain some sort of 'logic'
	// Will discuss later
	// Example usage:
	// To execute "select name from users where id = 3", call this function with:
	// readFromDB(*db, "users", "name", "id", "3")
	var result string

	rows, err := db.QueryRow("SELECT " + fields + " FROM " + table + " WHERE " + query + " = " + value).Scan(&result)

	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	return result
}
