package killbillApi

import "github.com/satori/go.uuid"

type User struct {
	UUID         uuid.UUID
	email        string
	firstname    string
	lastname     string
	transactions []Transaction
}

func CreateUser(email string, firstname string, lastname string) User {
	return User{
		uuid.NewV4(),
		email,
		firstname,
		lastname,
		make([]Transaction, 100),
	}
}
