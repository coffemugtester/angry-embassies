package config

import (
	"angry-embassies/conf"
	"data_persistance/client"
	"embassy_factory"
	"fmt"
)

type Dependencies struct {
	PersistorClient *client.Client
	MakerClient     *embassy_factory.Client
	//EmbassyService *services.EmbassyService
	//MgoService     *services.MgoService
	//GoogleService  *services.GoogleService
}

func InitDependencies() (Dependencies, error) {

	cfg := conf.LoadConfig()
	fmt.Printf("Config loaded: %v\n", cfg)

	persistorClient := client.NewClient(cfg.Mgo)
	fmt.Printf("PersistorClient created: %v\n", persistorClient)

	makerClient := embassy_factory.NewClient(cfg.ApiKey)

	//mgoUsecase := usecases.NewInsertUseCase(cfg.Mgo)
	//embassyUsecase := usecases.NewEmbassyUsecase(cfg.Domain)
	//googleUsecase := usecases.NewGoogleUsecase(cfg.ApiKey)

	return Dependencies{
		PersistorClient: persistorClient,
		MakerClient:     makerClient,
		//EmbassyService: services.NewEmbassyService(embassyUsecase),
		//MgoService:     services.NewMgoService(mgoUsecase),
		//GoogleService:  services.NewGoogleService(googleUsecase),
	}, nil
}
