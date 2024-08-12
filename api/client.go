package api

import (
	"angry-embassies/models"
	"api/maps"
)

type Client struct {
	*maps.Client
}

func NewClient(apiKey string) *Client {
	mapsClient := maps.NewMapsClient(apiKey)
	return &Client{
		mapsClient,
	}
}

func (c Client) GetGoogleID(placeQuery string) string {
	return c.Client.GetGoogleID(placeQuery)
}

func (c Client) GetPlaceDetails(placeQuery string) (models.PlaceDetails, error) {
	return c.Client.GetPlaceDetails(placeQuery)
}
