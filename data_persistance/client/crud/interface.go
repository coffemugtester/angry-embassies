package crud

import (
	"angry-embassies/conf"
	"embassy_factory"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoClient interface {
	Connect(mgoConf conf.MgoConfig) (*mongo.Client, error)
	Ping() error
	Collection(nameDB, nameCollection string) *mongo.Collection
	InsertOne(document embassy_factory.Embassy) (string, error)
}
