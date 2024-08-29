package services

type RepositoryService interface {
	InsertDocument(home, host string) (string, error)
}
