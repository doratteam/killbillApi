package model

import "github.com/satori/go.uuid"

type User struct {
	// User struct type
	UUID      uuid.UUID
	Email     string
	Firstname string
	Lastname  string
	Age       int
}
