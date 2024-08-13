package config

import (
	"angry-embassies/conf"
	"fmt"
	"repository/client"
	"repository/services"
	"repository/usecases"
)

type Dependencies struct {
	MgoService *services.MgoService
}

func InitDependencies() (Dependencies, error) {

	cfg := conf.LoadConfig()
	fmt.Printf("Config loaded: %v\n", cfg)

	persistorClient := client.NewClient(cfg.Mgo)
	fmt.Printf("PersistorClient created: %v\n", persistorClient)

	//makerClient := api.NewClient(cfg.ApiKey)

	mgoUseCase := *usecases.NewPersistenceUseCase(cfg.Mgo)

	return Dependencies{
		MgoService: services.NewMgoService(mgoUseCase),
	}, nil
}
