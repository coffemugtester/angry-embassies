package crud

import (
	"conf"
	"go.mongodb.org/mongo-driver/mongo"
	"models"
)

type MongoClient interface {
	Connect(mgoConf conf.MgoConfig) (*mongo.Client, error)
	Ping() error
	Collection(nameDB, nameCollection string) *mongo.Collection
	InsertOne(document models.Embassy) (string, error)
	FindOne(document models.Embassy) (models.Embassy, error)
	FindMany(document models.Embassy) ([]models.Embassy, error)
}
