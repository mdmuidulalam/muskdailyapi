package model

import "time"

type Account struct {
	FirstName      string
	LastName       string
	Email          string
	HashedPassword []byte
	Active         bool
	CreatedOn      time.Time
	ModifiedOn     *time.Time
	HashedCode     []byte
	CodeSentTime   time.Time
}
