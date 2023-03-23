package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Affiction/go-web-app/cmd/pkg/config"
	"github.com/Affiction/go-web-app/cmd/pkg/models"
)

var app *config.AppConfig

func NewTemplates(c *config.AppConfig) {
	app = c
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders the given template using the provided http.ResponseWriter.
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	// If we working in development mode, we create a cache for the each request
	if app.UseCache {
		// Create template cache
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// Get template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from cache")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	// Render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")

			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
