package data

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	model "muskdaily.com/model"
)

type AccountData struct {
	BaseData
}

func (this AccountData) InsertAccount(account model.Account) bool {
	collection := this.client.Database("muskdaily").Collection("accounts")

	_, err := collection.InsertOne(context.TODO(), account)

	if err != nil {
		panic(err)
	}

	return true
}

func (this AccountData) SelectAccounts(filter primitive.D) []*model.Account {
	var results []*model.Account
	collection := this.client.Database("muskdaily").Collection("accounts")

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	for cur.Next(context.TODO()) {
		var account model.Account
		err := cur.Decode(&account)
		if err != nil {
			panic(err)
		}

		results = append(results, &account)
	}

	return results
}

func (this AccountData) UpdateAccounts(filter, update primitive.D) (int64, int64) {
	collection := this.client.Database("muskdaily").Collection("accounts")
	updateResult, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	return updateResult.MatchedCount, updateResult.ModifiedCount
}
