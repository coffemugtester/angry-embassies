package main

import (
	config "configtest"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"repository/handlers"
)

func main() {

	deps, err := config.InitDependencies()
	if err != nil {
		fmt.Printf("config.InitDependencies error: %v\n", err)
		return
	}

	mgoService := deps.MgoService

	tmpl := template.Must(template.ParseFiles("templates/template.html", "templates/missions_overview.html"))
	r := gin.Default()
	r.SetHTMLTemplate(tmpl)

	//Serve static files (CSS, JS, images) from the static directory
	r.Static("/static", "./static")

	// Parse templates
	// TODO: wrap this in a InitApp function
	handler := handlers.NewEmbassyHandler(*mgoService)
	r.GET("/missions/:homeCountry/:hostCountry/:city", handler.GetDocument)
	r.GET("/missions/:homeCountry/:hostCountry", handler.GetDocuments)
	r.GET("/missions/:homeCountry", handler.GetDocuments)
	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	// Define the route to get country's embassies
	// TODO: item page - embassies/country
	// TODO: item page - embassies/country/country
	// TODO: item page - embassies/country/country/city

	r.Run()
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
