package client

import (
	"angry-embassies/conf"
	"angry-embassies/models"
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
