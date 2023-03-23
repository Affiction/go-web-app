package handlers

import (
	"net/http"

	"github.com/Affiction/go-web-app/cmd/pkg/render"
)

// Home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.template.tmpl")
}

// About page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.template.tmpl")
}
