package services

import (
	"models"
)

type RepositoryService interface {
	InsertDocument(embassy models.Embassy) (string, error)
	GetDocument(embassy models.Embassy) (models.Embassy, error)
}
