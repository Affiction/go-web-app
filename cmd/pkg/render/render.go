package render

import (
	"fmt"
	"html/template"
	"net/http"
)

var templateCache = make(map[string]*template.Template)

// RenderTemplate renders the given template using the provided http.ResponseWriter.
func RenderTemplate(w http.ResponseWriter, t string) {
	tmpl, ok := templateCache[t]
	if !ok {
		fmt.Println("Template - not cached")
		var err error
		tmpl, err = createTemplate(t)
		if err != nil {
			fmt.Printf("Error creating template %s: %v\n", t, err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		templateCache[t] = tmpl
	} else {
		fmt.Println("Template - cached")
	}

	err := tmpl.Execute(w, nil)
	if err != nil {
		fmt.Printf("Error rendering template %s: %v\n", t, err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// createTemplate creates and returns a new *template.Template instance using the given template name.
func createTemplate(tmpl string) (*template.Template, error) {
	paths := []string{
		fmt.Sprintf("./templates/%s", tmpl),
		"./templates/base.layout.tmpl",
	}
	return template.ParseFiles(paths...)
}
