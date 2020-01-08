package model

type Account struct {
	FirstName      string
	LastName       string
	Email          string
	HashedPassword []byte
	Active         bool
}
