package usecases

import (
	"angry_embassies/models"
	"api"
)

var _ GoogleMapsClient = (*EmbassyUsecase)(nil)

type EmbassyUsecase struct {
	ApiClient api.Client
}

func NewEmbassyUsecase(apiKey string) *EmbassyUsecase {
	mapsClient := api.NewClient(apiKey)
	return &EmbassyUsecase{
		ApiClient: *mapsClient,
	}
}

func (*EmbassyUsecase) GetEmbassy(gClient api.Client, home, host string) (models.Embassy, error) {
	var placeQuery string

	embassy := *models.NewEmbassy(home, host, false, "", "", "", models.PlaceDetails{})
	if embassy.IsConsulate {
		placeQuery = embassy.HomeCountry + " consulate in " + embassy.City + ", " + embassy.HostCountry
	} else {
		placeQuery = embassy.HomeCountry + " embassy in " + embassy.City + ", " + embassy.HostCountry
	}

	var err error
	embassy.GoogleID, err = gClient.GetGoogleID(placeQuery)
	if err != nil {
		return models.Embassy{}, err
	}

	return embassy, nil
}

func (e *EmbassyUsecase) GetEmbassyDetails(embassy models.Embassy) models.Embassy {

	var err error
	embassy.PlaceDetails, err = e.ApiClient.GetPlaceDetails(embassy.GoogleID)
	if err != nil {
		return models.Embassy{}
	}
	return embassy
}
