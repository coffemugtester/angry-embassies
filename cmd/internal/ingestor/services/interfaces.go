package services

import (
	"models"
)

type IngestionService interface {
	GetEmbassyDetails(embassy models.Embassy) (models.Embassy, error)
	GetEmbassy(embassy models.Embassy) (models.Embassy, error)
}
