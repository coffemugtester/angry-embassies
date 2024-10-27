package crud

import (
	"angry_embassies/conf"
	"angry_embassies/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ MongoClient = (*MongoImpl)(nil)

type MongoImpl struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func (t MongoImpl) Connect(mgoConf conf.MgoConfig) (*mongo.Client, error) {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mgoConf.MongoUri))
	if err != nil {
		fmt.Printf("Error connecting to MongoDB: %v", err)
		return nil, err
	}

	return client, nil
}

func (t MongoImpl) Ping() error {
	err := t.client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Printf("Error pinging MongoDB: %v", err)
		return err
	}
	return nil
}

func (t MongoImpl) Collection(nameDB, nameCollection string) *mongo.Collection {
	return t.client.Database(nameDB).Collection(nameCollection)
}

func (t MongoImpl) InsertOne(document models.Embassy) (string, error) {
	id, err := t.collection.InsertOne(context.TODO(), document)
	if err != nil {
		fmt.Printf("Error inserting document: %v", err)
		return "", err
	}

	return fmt.Sprintf("%v", id), nil
}

func (t MongoImpl) FindOne(document models.Embassy) (models.Embassy, error) {
	var embassy models.Embassy

	//The filter parameter must be a document containing query operators and can be used to select the document to be returned.
	//It cannot be nil. If the filter does not match any documents, a SingleResult with an error set to ErrNoDocuments will be returned.
	//If the filter matches multiple documents, one will be selected from the matched set.
	err := t.collection.FindOne(context.TODO(), document).Decode(&embassy)
	if err != nil {
		fmt.Printf("Error finding document: %v", err)
		return models.Embassy{}, err
	}

	return embassy, nil
}

type Client struct {
	mgoDB         string
	mgoCollection string
	mongoImpl     MongoClient
}

func NewCRUDClient(mgoConf conf.MgoConfig, mongoImpl MongoImpl) *Client {

	fmt.Printf("Creating new CRUD client with config: %v\n", mgoConf)

	mongoImpl.client, _ = mongoImpl.Connect(mgoConf)
	mongoImpl.collection = mongoImpl.Collection(mgoConf.MongoDb, mgoConf.MongoCollection)

	mongoImpl.Connect(mgoConf)
	fmt.Printf("Connected to MongoDB\n")
	fmt.Printf("t.client nil: %v\n", mongoImpl.client == nil)

	fmt.Printf("Pinging MongoDB\n")

	mongoImpl.Ping()
	return &Client{
		mgoDB:         mgoConf.MongoDb,
		mgoCollection: mgoConf.MongoCollection,
		mongoImpl:     mongoImpl,
	}
}
