package crud

import (
	"conf"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"models"
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

func (t MongoImpl) FindMany(document models.Embassy) ([]models.Embassy, error) {

	data := map[string]interface{}{
		"home_country": document.HomeCountry,
	}

	if document.HostCountry != "" {
		data["host_country"] = document.HostCountry
	}

	bsonData, err := bson.Marshal(data)
	if err != nil {
		fmt.Printf("Error marshalling data: %v\n", err)
		return nil, err
	}

	var filter bson.M
	err = bson.Unmarshal(bsonData, &filter)

	var embassy []models.Embassy

	curs, colError := t.collection.Find(context.TODO(), filter)
	if colError != nil {
		fmt.Printf("Error finding document: %v\n", err)
		return []models.Embassy{}, err
	}

	for curs.Next(context.Background()) {
		var emb models.Embassy
		err := curs.Decode(&emb)
		if err != nil {
			fmt.Printf("Error decoding document: %v\n", err)
			return []models.Embassy{}, err
		}
		embassy = append(embassy, emb)
	}
	fmt.Println("Embassy found: ", embassy)

	return embassy, nil
}

func (t MongoImpl) FindOne(document models.Embassy) (models.Embassy, error) {

	filter := bson.M{
		"home_country": document.HomeCountry,
		"host_country": document.HostCountry,
		"city":         document.City,
	}

	var embassy models.Embassy
	err := t.collection.FindOne(context.TODO(), filter).Decode(&embassy)
	if err != nil {
		fmt.Printf("Error finding document: %v\n", err)
		return models.Embassy{}, err
	}

	fmt.Println("Embassy found: ", embassy)
	fmt.Println("Embassy place details: ", *embassy.PlaceDetails)

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
