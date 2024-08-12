package embassy_factory

import (
	"embassy_factory/mapz"
)

type Client struct {
	*mapz.Client
}

func NewClient(apiKey string) *Client {
	mapzClient := mapz.NewMapzClient(apiKey)
	return &Client{
		mapzClient,
	}
}

func (c Client) GetGoogleID(placeQuery string) string {
	return c.Client.GetGoogleID(placeQuery)
}

func (c Client) GetPlaceDetails(placeQuery string) (PlaceDetails, error) {
	return c.Client.GetPlaceDetails(placeQuery)
}
