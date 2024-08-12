package usecases

import "embassy_factory"

type PersistenceClient interface {
	InsertDocument(document embassy_factory.Embassy) (string, error)
}
