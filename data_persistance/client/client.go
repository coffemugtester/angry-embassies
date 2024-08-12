package client

import (
	"angry-embassies/conf"
	"data_persistance/client/crud"
	"embassy_factory"
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

func (c Client) InsertDocument(document embassy_factory.Embassy) (string, error) {
	return crud.InsertDocument(c.Client, document)
}
