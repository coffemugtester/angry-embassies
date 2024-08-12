package usecases

import "data_persistance/client"

var _ PersistenceClient = (*PersistenceUsecase)(nil)

type PersistenceUsecase struct {
	client.Client
}
