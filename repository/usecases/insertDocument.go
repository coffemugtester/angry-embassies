package usecases

import "repository/client"

var _ PersistenceClient = (*PersistenceUsecase)(nil)

type PersistenceUsecase struct {
	client.Client
}
