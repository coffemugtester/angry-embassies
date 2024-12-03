package crud

import (
	"fmt"
	"models"
)

func InsertDocument(client *Client, document models.Embassy) (string, error) {
	fmt.Printf("Inserting document: %v\n", document)

	insertedDoc, err := client.mongoImpl.InsertOne(document)
	if err != nil {
		return "", err
	}

	return insertedDoc, nil
}

// TODO: Implement GetDocument

func GetDocument(client *Client, document models.Embassy) (models.Embassy, error) {

	retrievedDoc, err := client.mongoImpl.FindOne(document)
	if err != nil {
		return models.Embassy{}, err
	}

	return retrievedDoc, nil
}

func GetDocuments(client *Client, document models.Embassy) ([]models.Embassy, error) {

	retrievedDocs, err := client.mongoImpl.FindMany(document)
	if err != nil {
		return []models.Embassy{}, err
	}

	return retrievedDocs, nil
}
