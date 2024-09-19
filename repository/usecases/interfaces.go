package usecases

import (
	"angry_embassies/models"
)

type PersistenceClient interface {
	InsertDocument(document models.Embassy) (string, error)
}
