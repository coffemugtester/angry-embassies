package services

import (
	"angry-embassies/models"
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

func (m *MgoService) InsertDocument(home, host string) (string, error) {

	embassy := models.Embassy{
		HomeCountry: home,
		HostCountry: host,
	}

	return m.useCase.InsertDocument(embassy)
}
