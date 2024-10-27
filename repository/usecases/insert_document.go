package usecases

import (
	"angry_embassies/conf"
	"angry_embassies/models"
	"repository/client"
)

var _ PersistenceClient = (*PersistenceUseCase)(nil)

type PersistenceUseCase struct {
	client client.Client
}

func NewPersistenceUseCase(mgoConf conf.MgoConfig) *PersistenceUseCase {
	client := *client.NewClient(mgoConf)
	return &PersistenceUseCase{
		client,
	}
}

func (p *PersistenceUseCase) InsertDocument(document models.Embassy) (string, error) {
	return p.client.InsertDocument(document)
}

func (p *PersistenceUseCase) GetDocument(document models.Embassy) (models.Embassy, error) {
	return p.client.GetDocument(document)
}
