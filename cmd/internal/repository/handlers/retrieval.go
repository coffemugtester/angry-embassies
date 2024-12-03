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

func (e *EmbassyHandler) GetDocument(c *gin.Context) {

	var embassy models.Embassy
	if c.Param("homeCountry") != "" {
		embassy.HomeCountry = c.Param("homeCountry")
	}
	if c.Param("hostCountry") != "" {
		embassy.HostCountry = c.Param("hostCountry")
	}
	if c.Param("city") != "" {
		embassy.City = c.Param("city")
	}

	// Call the service to get the embassies
	mgoResult, err := e.service.GetDocument(embassy)
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

func (e *EmbassyHandler) GetDocuments(c *gin.Context) {

	var embassy models.Embassy
	if c.Param("homeCountry") != "" {
		embassy.HomeCountry = c.Param("homeCountry")
	}
	if c.Param("hostCountry") != "" {
		embassy.HostCountry = c.Param("hostCountry")
	}

	// Call the service to get the embassies
	mgoResult, err := e.service.GetDocuments(embassy)
	if err != nil {
		c.JSON(500, gin.H{
			// TODO: not exposing the error message to the user
			"error": fmt.Sprint(err),
		})
		return
	}

	//tmpl.ExecuteTemplate(w, "home.html", nil) instead of the line below
	c.HTML(200, "missions_overview.html", mgoResult)
}
