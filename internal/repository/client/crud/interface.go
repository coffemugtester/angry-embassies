package crud

import (
	"conftest"
	"go.mongodb.org/mongo-driver/mongo"
	"models"
)

type MongoClient interface {
	Connect(mgoConf conftest.MgoConfig) (*mongo.Client, error)
	Ping() error
	Collection(nameDB, nameCollection string) *mongo.Collection
	InsertOne(document models.Embassy) (string, error)
}
