package d7mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	options *setupOptions
	Client  *mongo.Client
}

func NewMongoClient(opts ...SetupOption) *MongoClient {
	opt := mergeSetupOptions(opts...)

	// connect
	client, err := mongo.Connect(context.TODO(),
		options.Client().ApplyURI(opt.getUri()).SetAuth(opt.getAuth()))
	if err != nil {
		panic(err)
	}

	// return client
	return &MongoClient{
		options: opt,
		Client:  client,
	}
}
