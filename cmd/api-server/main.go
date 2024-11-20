package main

import (
	"conf"
	"fmt"
	"github.com/gin-gonic/gin"
	"models"
	"repository/services"
	"repository/usecases"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

// TODO: set up basic routing
// TODO: host country's specific embassies - /country/embassies
// TODO: host country's specific embassy - /country/embassies/country/city
// TODO: home country's missions - /country/missions

func main() {

	con := conf.LoadConfig()
	mgoUseCase := usecases.NewPersistenceUseCase(con.Mgo)
	mgoService := services.NewMgoService(*mgoUseCase)

	r := gin.Default()
	// TODO: item page - missions/country
	// TODO: item page - missions/country/country
	r.GET("/missions/:homeCountry/:hostCountry", func(c *gin.Context) {
		homeCountry := c.Param("homeCountry")
		hostCountry := c.Param("hostCountry")

		// Call the service to get the embassies
		// TODO: add GetDocuments method to get all embassies for a country
		mgoResult, err := mgoService.GetDocument(models.Embassy{
			HomeCountry: homeCountry,
			HostCountry: hostCountry,
		})
		if err != nil {
			c.JSON(500, gin.H{
				// TODO: not exposing the error message to the user
				"error": fmt.Sprint(err),
			})
			return
		}

		c.JSON(200, gin.H{
			"country": mgoResult,
		})
	})

	// TODO: item page - missions/country/country/city
	// Define the route to get country's embassies
	r.GET("/missions/:homeCountry/:hostCountry/:city", func(c *gin.Context) {
		homeCountry := c.Param("homeCountry")
		hostCountry := c.Param("hostCountry")
		city := c.Param("city")

		// Call the service to get the embassies
		mgoResult, err := mgoService.GetDocument(models.Embassy{
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

		c.JSON(200, gin.H{
			"country": mgoResult,
		})
	})
	// TODO: item page - embassies/country
	// TODO: item page - embassies/country/country
	// TODO: item page - embassies/country/country/city

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.Run()
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
