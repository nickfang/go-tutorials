package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	// Parse the template file
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	// Execute the template
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	_, inMap := tc[t]
	fmt.Println("Checking if template is in cache", t)
	if !inMap {
		fmt.Println("creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			fmt.Println("error creating template cache:", err)
			return
		}
	} else {
		fmt.Println("template already in cache")
	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	tc[t] = tmpl

	return nil
}
