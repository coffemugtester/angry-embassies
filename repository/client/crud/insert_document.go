package crud

import (
	"angry_embassies/models"
	"fmt"
)

func InsertDocument(client *Client, document models.Embassy) (string, error) {
	fmt.Printf("Inserting document: %v\n", document)

	client.mongoImpl.Collection(client.mgoDB, client.mgoCollection)

	insertedDoc, err := client.mongoImpl.InsertOne(document)
	if err != nil {
		return "", err
	}

	return insertedDoc, nil
}

func GetDocument(client *Client, document models.Embassy) (models.Embassy, error) {
	fmt.Printf("Getting document: %v\n", document)

	client.mongoImpl.Collection(client.mgoDB, client.mgoCollection)

	return models.Embassy{}, nil
	//doc, err := client.mongoImpl.FindOne(document
}
