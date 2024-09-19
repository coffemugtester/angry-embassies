package services

import (
	"angry_embassies/models"
	"embassy_sync/usecases"
)

var _ EmbassySyncService = (*EmbassyService)(nil)

type EmbassyService struct {
	gglUsecase usecases.EmbassyUsecase
}

func NewEmbassyService(gglUsecase usecases.EmbassyUsecase) *EmbassyService {
	return &EmbassyService{
		gglUsecase: gglUsecase,
	}
}
func (e *EmbassyService) GetEmbassyDetails(home, host string) (models.Embassy, error) {
	embassy, err := e.GetEmbassy(home, host)
	if err != nil {
		return models.Embassy{}, err
	}
	return e.gglUsecase.GetEmbassyDetails(embassy), nil
}

func (e *EmbassyService) GetEmbassy(home, host string) (models.Embassy, error) {
	return e.gglUsecase.GetEmbassy(e.gglUsecase.ApiClient, home, host)
}
