package usecases

import (
	"models"
)

type PersistenceClient interface {
	InsertDocument(document models.Embassy) (string, error)
	GetDocument(document models.Embassy) (models.Embassy, error)
	GetDocuments(document models.Embassy) ([]models.Embassy, error)
	// TODO: add GetDocuments method
}
