package main

import (
	"fmt"
	"log"
	"net/http"
	"web-app/pkg/config"
	"web-app/pkg/handlers"
	"web-app/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", repo.Home)
	http.HandleFunc("/about", repo.About)
	http.HandleFunc("/favicon.ico", handlers.Favicon)

	fmt.Printf("Starting application on port %s\n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
