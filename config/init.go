package config

import (
	"angry_embassies/conf"
	embs "embassy_sync/services"
	emb "embassy_sync/usecases"
	"fmt"
	"repository/client"
	"repository/services"
	"repository/usecases"
)

type Dependencies struct {
	MgoService *services.MgoService
	GglService *embs.EmbassyService
}

func InitDependencies() (Dependencies, error) {

	cfg := conf.LoadConfig()
	fmt.Printf("Config loaded: %v\n", cfg)

	persistorClient := client.NewClient(cfg.Mgo)
	fmt.Printf("PersistorClient created: %v\n", persistorClient)

	mgoUseCase := *usecases.NewPersistenceUseCase(cfg.Mgo)
	gglUseCase := emb.NewEmbassyUsecase(cfg.ApiKey)
	fmt.Printf("ggUseCase created: %v\n", gglUseCase)

	//apiClient := api.NewClient(cfg.ApiKey)

	return Dependencies{
		MgoService: services.NewMgoService(mgoUseCase),
		GglService: embs.NewEmbassyService(*gglUseCase),
		//ApiClient: apiClient,
	}, nil
}
