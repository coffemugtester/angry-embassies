package services

import (
	"models"
)

type EmbassySyncService interface {
	GetEmbassyDetails(embassy models.Embassy) (models.Embassy, error)
	GetEmbassy(embassy models.Embassy) (models.Embassy, error)
}
