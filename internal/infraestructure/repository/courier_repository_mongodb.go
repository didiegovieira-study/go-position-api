package repository

import (
	"context"
	"fmt"

	"github.com/didiegovieira/go-position-api/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CourierMongodb struct {
	Client     *mongo.Client
	collection *mongo.Collection
}

func NewCourierMongodb(client *mongo.Client) *CourierMongodb {
	benchmarkMongoDB := &CourierMongodb{
		Client:     client,
		collection: client.Database("courier").Collection("courier"),
	}

	benchmarkMongoDB.createIndexes()

	return benchmarkMongoDB
}

func (c CourierMongodb) createIndexes() error {
	_, err := c.collection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{Keys: bson.D{{Key: "_id", Value: 1}}},
	}, &options.CreateIndexesOptions{})

	if err != nil {
		return err
	}

	return nil
}

func (c *CourierMongodb) Save(courier *entity.Courier) error {
	filter := bson.M{"_id": courier.ID}

	_, err := c.collection.ReplaceOne(
		context.Background(), filter, courier, options.Replace().SetUpsert(true),
	)

	if err != nil {
		return err
	}

	return nil
}

func (c *CourierMongodb) GetById(id string) (*entity.Courier, error) {
	var courier entity.Courier

	err := c.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&courier)
	if err != nil {
		return nil, err
	}

	return &courier, nil
}

func (c *CourierMongodb) GetNearestCouriers(lat, lng float64) {
	mongoDBHQ := bson.D{{"type", "Point"}, {"coordinates", []float64{lng, lat}}}
	filter := bson.D{
		{"actual_position", bson.D{
			{"$near", bson.D{
				{"$geometry", mongoDBHQ},
				{"$maxDistance", 1000},
			}},
		}},
	}

	var couriers []entity.Courier
	output, err := c.collection.Find(context.TODO(), filter)
	if err = output.All(context.TODO(), &couriers); err != nil {
		panic(err)
	}

	for _, courier := range couriers {
		res, _ := bson.MarshalExtJSON(courier, false, false)
		fmt.Println(string(res))
	}
}
