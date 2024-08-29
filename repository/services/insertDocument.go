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

	embassy := *models.NewEmbassy(apiClient, home, host, false, "", "", "", models.PlaceDetails{})
	//TODO: build placeQuery (embassy/consulate of home in city host)
	embassy.ApiClient.GetGoogleID(embassy.HomeCountry)
	//TODO: make GetPlaceDetails void
	embassy.ApiClient.GetPlaceDetails(embassy.GoogleID)

	return m.useCase.InsertDocument(embassy)
}
