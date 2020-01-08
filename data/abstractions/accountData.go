package data

import (
	model "muskdaily.com/model"
)

type AccountData interface {
	Connect()
	Disconnect()
	AddAccount(account model.Account) bool
}
