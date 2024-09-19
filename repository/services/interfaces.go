package services

import "angry_embassies/models"

type RepositoryService interface {
	InsertDocument(embassy models.Embassy) (string, error)
}
