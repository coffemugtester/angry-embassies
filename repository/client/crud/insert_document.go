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
