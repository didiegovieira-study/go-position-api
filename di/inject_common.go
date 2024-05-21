package di

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func provideMongoDbClient() (*mongo.Client, func(), error) {
	mongodbClient, err := mongo.Connect(
		context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/"),
	)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		_ = mongodbClient.Disconnect(context.Background())
	}

	return mongodbClient, cleanup, nil
}
