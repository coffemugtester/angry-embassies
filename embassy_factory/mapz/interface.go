package mapz

import "maker"

type MapsClient interface {
	// GetGoogleID returns the Google id of the location
	GetGoogleID(placeQuery string) string
	GetPlaceDetails(placeQuery string) (embassy_factory.PlaceDetails, error)
}
