package maps

import "angry-embassies/models"

type ClientInterface interface {
	// GetGoogleID returns the Google id of the location
	GetGoogleID(placeQuery string) string
	GetPlaceDetails(placeQuery string) (models.PlaceDetails, error)
}
