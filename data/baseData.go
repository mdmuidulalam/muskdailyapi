package data

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	config "muskdaily.com/config"
)

type BaseData struct {
	client *mongo.Client
}

func (this *BaseData) Connect() {
	configuration := config.GetConfiguration()

	var err error
	this.client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(configuration.Database.Path+":"+configuration.Database.Port))

	if err != nil {
		panic(err)
	}

	err = this.client.Ping(context.TODO(), nil)

	if err != nil {
		panic(err)
	}
}

func (this *BaseData) Disconnect() {
	err := this.client.Disconnect(context.TODO())

	if err != nil {
		panic(err)
	}
}
