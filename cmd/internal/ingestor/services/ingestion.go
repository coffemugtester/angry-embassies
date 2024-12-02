package services

import (
	"ingestor/usecases"
	"models"
)

var _ IngestionServiceImpl = (*IngestionService)(nil)

type IngestionService struct {
	gglUsecase usecases.EmbassyUsecase
}

func NewIngestionService(gglUsecase usecases.EmbassyUsecase) *IngestionService {
	return &IngestionService{
		gglUsecase: gglUsecase,
	}
}

func (e *IngestionService) GetEmbassyDetails(embassy models.Embassy) (models.Embassy, error) {
	embassy, err := e.GetEmbassy(embassy)
	if err != nil {
		return models.Embassy{}, err
	}
	return e.gglUsecase.GetEmbassyDetails(embassy), nil
}

func (e *IngestionService) GetEmbassy(embassy models.Embassy) (models.Embassy, error) {
	return e.gglUsecase.GetEmbassy(e.gglUsecase.ApiClient, embassy)
}
