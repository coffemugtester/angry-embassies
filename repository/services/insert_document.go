package services

import (
	"angry_embassies/models"
	"repository/usecases"
)

var _ RepositoryService = (*MgoService)(nil)

type MgoService struct {
	useCase usecases.PersistenceUseCase
}

func NewMgoService(useCase usecases.PersistenceUseCase) *MgoService {
	return &MgoService{
		useCase,
	}
}

func (m *MgoService) InsertDocument(embassy models.Embassy) (string, error) {
	return m.useCase.InsertDocument(embassy)
}
