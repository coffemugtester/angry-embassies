package usecases

import (
	"models"
)

type PersistenceClient interface {
	InsertDocument(document models.Embassy) (string, error)
}
