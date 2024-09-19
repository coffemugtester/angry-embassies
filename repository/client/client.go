package client

import (
	"angry_embassies/conf"
	"angry_embassies/models"
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
