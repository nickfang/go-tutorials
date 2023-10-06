package main

import (
	"github.com/bmizerany/pat"
	"net/http"
	"web-app/pkg/config"
	"web-app/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux.Get("/favicon.ico", http.HandlerFunc(handlers.Repo.Favicon))

	return mux
}
