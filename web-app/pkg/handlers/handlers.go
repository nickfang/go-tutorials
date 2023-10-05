package handlers

import (
	"fmt"
	"net/http"
	"web-app/pkg/render"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home page handler called")
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}

func Favicon(w http.ResponseWriter, r *http.Request) {
	// do nothing
}
