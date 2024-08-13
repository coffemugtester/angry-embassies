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

func (m *MgoService) InsertDocument(document models.Embassy) (string, error) {
<<<<<<< HEAD
=======

	//TODO: the repository service creates an embassy object

>>>>>>> 87c1cb1 (add repositiory's usecase and service)
	return m.useCase.InsertDocument(document)
}
