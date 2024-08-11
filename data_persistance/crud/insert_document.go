package crud

import (
	"fmt"
)

func InsertDocument(client *Client, document string) (string, error) {
	fmt.Printf("Inserting document: %v\n", document)

	client.mongoImpl.Collection(client.mgoDB, client.mgoCollection)

	insertedDoc, err := client.mongoImpl.InsertOne(document)
	if err != nil {
		return "", err
	}

	return insertedDoc, nil
}
