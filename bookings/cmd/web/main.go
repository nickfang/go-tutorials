package main

import (
	"fmt"
	"github.com/nickfang/bookings/internal/config"
	"github.com/nickfang/bookings/internal/handlers"
	"github.com/nickfang/bookings/internal/render"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour        // 24 hours in seconds
	session.Cookie.Persist = true            // persist session across browser restarts
	session.Cookie.Secure = app.InProduction // set to true in production
	session.Cookie.SameSite = http.SameSiteLaxMode

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false
	app.Session = session

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Printf("Starting application on port %s\n", portNumber)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
