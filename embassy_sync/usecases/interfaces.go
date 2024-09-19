package usecases

import (
	"angry_embassies/models"
	"api"
)

type GoogleMapsClient interface {
	GetEmbassy(gClient api.Client, home, host string) (models.Embassy, error)
	GetEmbassyDetails(embassy models.Embassy) models.Embassy
}
