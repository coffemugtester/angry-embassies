package services

import "angry-embassies/models"

type RepositoryService interface {
	InsertDocument(document models.Embassy) (string, error)
}
