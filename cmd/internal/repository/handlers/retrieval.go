package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"models"
	"repository/services"
)

var _ EmbassyHandlerImpl = (*EmbassyHandler)(nil)

type EmbassyHandler struct {
	service services.MgoService
}

func NewEmbassyHandler(service services.MgoService) *EmbassyHandler {
	return &EmbassyHandler{
		service: service,
	}
}

func (e *EmbassyHandler) GetDocuments(c *gin.Context) {
	homeCountry := c.Param("homeCountry")
	hostCountry := c.Param("hostCountry")
	city := c.Param("city")

	// Call the service to get the embassies
	mgoResult, err := e.service.GetDocuments(models.Embassy{
		HomeCountry: homeCountry,
		HostCountry: hostCountry,
		City:        city,
	})
	if err != nil {
		c.JSON(500, gin.H{
			// TODO: not exposing the error message to the user
			"error": fmt.Sprint(err),
		})
		return
	}

	//tmpl.ExecuteTemplate(w, "home.html", nil) instead of the line below
	c.HTML(200, "template.html", mgoResult)

}

//func parseParameters(c *gin.Context) (models.Embassy, error) {
//
//	homeCountry := c.Param("homeCountry")
//	hostCountry := c.Param("hostCountry")
//	city := c.Param("city")
//
//	filter := make(map[string]interface{})
//	if homeCountry != "" {
//		filter["home_country"] = homeCountry
//	}
//	if hostCountry != "" {
//		filter["host_country"] = hostCountry
//	}
//	if city != "" {
//		filter["city"] = city
//	}
//
//	serializedParams, err := json.Marshal(filter)
//	if err != nil {
//		return models.Embassy{}, err
//	}
//
//	embassy := models.Embassy{}
//	err = json.Unmarshal(serializedParams, &embassy)
//	if err != nil {
//		return models.Embassy{}, err
//	}
//	return embassy, nil
//}
