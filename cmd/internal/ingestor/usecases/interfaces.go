package usecases

import (
	"api"
	"models"
)

type GoogleMapsClient interface {
	GetEmbassy(gClient api.Client, embassy models.Embassy) (models.Embassy, error)
	GetEmbassyDetails(embassy models.Embassy) models.Embassy
}
