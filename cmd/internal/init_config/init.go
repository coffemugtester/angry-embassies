package init_config

import (
	"conf"
	"fmt"
	"github.com/gin-gonic/gin"
	embs "ingestor/services"
	emb "ingestor/usecases"
	"repository/client"
	"repository/services"
	"repository/usecases"
)

type Dependencies struct {
	MgoService       *services.MgoService
	IngestionService *embs.IngestionService
}

type Handlers struct {
	GetEmbassy gin.HandlerFunc
}

func InitHandlers(deps Dependencies) Handlers {
	// TODO: repositoryHandler := repository.NewHandler(deps.MgoService)
	return Handlers{
		// TODO: each endpoint gets one specific method defined in the handler
		GetEmbassy: func(c *gin.Context) {},
	}
}

func InitDependencies() (Dependencies, error) {

	cfg := conf.LoadConfig()
	fmt.Printf("Config loaded: %v\n", cfg)

	persistorClient := client.NewClient(cfg.Mgo)
	fmt.Printf("PersistorClient created: %v\n", persistorClient)

	mgoUseCase := *usecases.NewPersistenceUseCase(cfg.Mgo)
	gglUseCase := emb.NewEmbassyUsecase(cfg.ApiKey)
	fmt.Printf("ggUseCase created: %v\n", gglUseCase)

	return Dependencies{
		MgoService:       services.NewMgoService(mgoUseCase),
		IngestionService: embs.NewIngestionService(*gglUseCase),
	}, nil
}
