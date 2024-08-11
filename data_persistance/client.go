package data_persistance

import (
	"angry-embassies/conf"
	"data_persistance/crud"
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

func (c Client) InsertDocument(document string) (string, error) {
	return crud.InsertDocument(c.Client, document)
}
