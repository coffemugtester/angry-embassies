package main

import (
	config "configtest"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"models"
	"repository/handlers"
)

// TODO: load pinture to mongo too
func main() {

	deps, err := config.InitDependencies()
	if err != nil {
		fmt.Printf("config.InitDependencies error: %v\n", err)
		return
	}

	mgoService := deps.MgoService

	r := gin.Default()

	//Serve static files (CSS, JS, images) from the static directory
	r.Static("/static", "./static")

	//// Parse multiple templates
	//tmpl := template.Must(template.ParseFiles(
	//	"templates/base.html",
	//	"templates/home.html",
	//	"templates/about.html",
	//))

	// Parse templates
	tmpl := template.Must(template.ParseFiles("templates/template.html"))
	// TODO: how to handle more than one template
	r.SetHTMLTemplate(tmpl)

	handler := handlers.NewEmbassyHandler(*mgoService)
	r.GET("/htmltest/:homeCountry/:hostCountry/:city", handler.GetDocuments)

	r.GET("/missions/:homeCountry", func(c *gin.Context) {
		homeCountry := c.Param("homeCountry")

		// Call the service to get the embassies
		// TODO: add GetDocuments method to get all embassies for a country
		mgoResult, err := mgoService.GetDocuments(models.Embassy{
			HomeCountry: homeCountry,
		})
		if err != nil {
			c.JSON(500, gin.H{
				// TODO: not exposing the error message to the user
				"error": fmt.Sprint(err),
			})
			return
		}

		c.JSON(200, gin.H{
			"countries": mgoResult,
		})
	})
	r.GET("/missions/:homeCountry/:hostCountry", func(c *gin.Context) {
		homeCountry := c.Param("homeCountry")
		hostCountry := c.Param("hostCountry")

		// Call the service to get the embassies
		mgoResult, err := mgoService.GetDocuments(models.Embassy{
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
			"countries": mgoResult,
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
