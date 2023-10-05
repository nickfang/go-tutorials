package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// Parse the template file
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	// Execute the template
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}
