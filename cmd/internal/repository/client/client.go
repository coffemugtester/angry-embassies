package client

import (
	"conf"
	"models"
	"repository/client/crud"
)

type Client struct {
	*crud.Client
}

func NewClient(mgo conf.MgoConfig) *Client {
	crudClient := crud.NewCRUDClient(mgo, crud.MongoImpl{})
	return &Client{
		crudClient,
	}
}

func (c Client) InsertDocument(document models.Embassy) (string, error) {
	return crud.InsertDocument(c.Client, document)
}

func (c Client) GetDocument(document models.Embassy) (models.Embassy, error) {
	return crud.GetDocument(c.Client, document)
}
