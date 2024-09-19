package crud

import (
	"angry_embassies/conf"
	"angry_embassies/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoClient interface {
	Connect(mgoConf conf.MgoConfig) (*mongo.Client, error)
	Ping() error
	Collection(nameDB, nameCollection string) *mongo.Collection
	InsertOne(document models.Embassy) (string, error)
}
