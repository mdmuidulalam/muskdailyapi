package data

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	config "muskdaily.com/config"
)

type Data struct {
	client        *mongo.Client
	Configuration *config.Configuration
}

func (this *Data) Connect() {
	var err error
	this.client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(this.Configuration.Database.Path+":"+this.Configuration.Database.Port))

	if err != nil {
		panic(err)
	}
}

func (this *Data) Disconnect() {
	err := this.client.Disconnect(context.TODO())

	if err != nil {
		panic(err)
	}
}
