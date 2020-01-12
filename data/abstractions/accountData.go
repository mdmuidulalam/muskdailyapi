package data

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	model "muskdaily.com/model"
)

type AccountData interface {
	Connect()
	Disconnect()
	SelectAccounts(filter primitive.D) []*model.Account
	InsertAccount(account model.Account) bool
	UpdateAccounts(filter, update primitive.D) (int64, int64)
}
