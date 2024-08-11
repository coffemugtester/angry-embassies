package crud

import (
	"angry-embassies/conf"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoClient interface {
	Connect(mgoConf conf.MgoConfig) (*mongo.Client, error)
	Ping() error
	Collection(nameDB, nameCollection string) *mongo.Collection
	InsertOne(document string) (string, error)
}
