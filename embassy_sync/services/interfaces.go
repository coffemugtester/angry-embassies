package services

import "angry_embassies/models"

type EmbassySyncService interface {
	GetEmbassyDetails(home, host string) (models.Embassy, error)
	GetEmbassy(home, host string) (models.Embassy, error)
}
