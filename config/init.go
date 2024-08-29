package config

import (
	"angry-embassies/conf"
	"api"
	"fmt"
	"repository/client"
	"repository/services"
	"repository/usecases"
)

type Dependencies struct {
	MgoService *services.MgoService
	ApiClient  *api.Client
}

func InitDependencies() (Dependencies, error) {

	cfg := conf.LoadConfig()
	fmt.Printf("Config loaded: %v\n", cfg)

	persistorClient := client.NewClient(cfg.Mgo)
	fmt.Printf("PersistorClient created: %v\n", persistorClient)

	mgoUseCase := *usecases.NewPersistenceUseCase(cfg.Mgo)
	apiClient := api.NewClient(cfg.ApiKey)

	return Dependencies{
		MgoService: services.NewMgoService(mgoUseCase),
		ApiClient:  apiClient,
	}, nil
}
