package data

import (
	"context"

	model "muskdaily.com/model"
)

type AccountData struct {
	BaseData
}

func (this AccountData) AddAccount(account model.Account) bool {
	collection := this.client.Database("muskdaily").Collection("accounts")

	_, err := collection.InsertOne(context.TODO(), account)

	if err != nil {
		panic(err)
	}

	return true
}
