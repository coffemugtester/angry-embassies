package usecases

import "embassy_factory"

type PersistenceClient interface {
	InsertDocument(document api.Embassy) (string, error)
}
