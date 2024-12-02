package usecases

import (
	"api"
	"fmt"
	"models"
	"net/url"
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

func (*EmbassyUsecase) GetEmbassy(gClient api.Client, embassy models.Embassy) (models.Embassy, error) {
	var placeQuery string

	//TODO: pass complete embassy
	if embassy.IsConsulate {
		placeQuery = embassy.HomeCountry + " consulate in " + embassy.City + ", " + embassy.HostCountry
	} else {
		placeQuery = embassy.HomeCountry + " embassy in " + embassy.City + ", " + embassy.HostCountry
	}

	placeQuery = url.QueryEscape(placeQuery)

	var err error
	embassy.GoogleID, err = gClient.GetGoogleID(placeQuery)
	//TODO: handle encoding
	embassy.MapLink = fmt.Sprintf("https://www.google.com/maps?q=%s", placeQuery)
	fmt.Println("Map link: ", embassy.MapLink)

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
	embassy.Picture, err = e.ApiClient.GetPlacePicture(embassy.PlaceDetails.Result.Photos[0].PhotoReference)
	if err != nil {
		return models.Embassy{}
	}
	return embassy
}
