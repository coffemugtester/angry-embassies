package services

import "api"

type RepositoryService interface {
	InsertDocument(apiClient api.Client, home, host string) (string, error)
}
