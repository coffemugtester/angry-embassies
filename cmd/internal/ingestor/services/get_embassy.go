package services

import (
	"embassy_sync/usecases"
	"models"
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
func (e *EmbassyService) GetEmbassyDetails(embassy models.Embassy) (models.Embassy, error) {
	embassy, err := e.GetEmbassy(embassy)
	if err != nil {
		return models.Embassy{}, err
	}
	return e.gglUsecase.GetEmbassyDetails(embassy), nil
}

func (e *EmbassyService) GetEmbassy(embassy models.Embassy) (models.Embassy, error) {
	return e.gglUsecase.GetEmbassy(e.gglUsecase.ApiClient, embassy)
}
