package services

import (
	"angry-embassies/models"
	"api"
	"repository/usecases"
)

var _ RepositoryService = &MgoService{}

type MgoService struct {
	useCase usecases.PersistenceUseCase
}

func NewMgoService(useCase usecases.PersistenceUseCase) *MgoService {
	return &MgoService{
		useCase,
	}
}

func (m *MgoService) InsertDocument(apiClient api.Client, home, host string) (string, error) {
	var placeQuery string

	embassy := *models.NewEmbassy(apiClient, home, host, false, "", "", "", models.PlaceDetails{})
	if embassy.IsConsulate {
		placeQuery = embassy.HomeCountry + " consulate in " + embassy.City + ", " + embassy.HostCountry
	} else {
		placeQuery = embassy.HomeCountry + " embassy in " + embassy.City + ", " + embassy.HostCountry
	}

	embassy.ApiClient.GetGoogleID(placeQuery)
	//TODO: make GetPlaceDetails void
	placeDetails, err := embassy.ApiClient.GetPlaceDetails(embassy.GoogleID)
	if err != nil {
		return "", err
	}

	embassy.PlaceDetails = placeDetails

	return m.useCase.InsertDocument(embassy)
}
