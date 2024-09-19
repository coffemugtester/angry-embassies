package maps

import "angry_embassies/models"

type ClientInterface interface {
	// GetGoogleID returns the Google id of the location
	GetGoogleID(placeQuery string) (string, error)
	GetPlaceDetails(placeQuery string) (*models.PlaceDetails, error)
}
