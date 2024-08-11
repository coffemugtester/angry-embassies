package config

import (
	"angry-embassies/conf"
	"data_persistance"
	"fmt"
)

type Dependencies struct {
	Client *data_persistance.Client
	//EmbassyService *services.EmbassyService
	//MgoService     *services.MgoService
	//GoogleService  *services.GoogleService
}

func InitDependencies() (Dependencies, error) {

	cfg := conf.LoadConfig()
	fmt.Printf("Config loaded: %v\n", cfg)

	client := data_persistance.NewClient(cfg.Mgo)
	fmt.Printf("Client created: %v\n", client)

	//mgoUsecase := usecases.NewInsertUseCase(cfg.Mgo)
	//embassyUsecase := usecases.NewEmbassyUsecase(cfg.Domain)
	//googleUsecase := usecases.NewGoogleUsecase(cfg.ApiKey)

	return Dependencies{
		Client: client,
		//EmbassyService: services.NewEmbassyService(embassyUsecase),
		//MgoService:     services.NewMgoService(mgoUsecase),
		//GoogleService:  services.NewGoogleService(googleUsecase),
	}, nil
}
