package killbillApi

import "github.com/satori/go.uuid"

type User struct {
	UUID         uuid.UUID
	owner        string
	transactions []Transaction
}
