package services

import (
	"models"
)

type IngestionServiceImpl interface {
	GetEmbassyDetails(embassy models.Embassy) (models.Embassy, error)
	GetEmbassy(embassy models.Embassy) (models.Embassy, error)
}
