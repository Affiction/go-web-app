package handlers

import (
	"net/http"

	"github.com/Affiction/go-web-app/cmd/pkg/config"
	"github.com/Affiction/go-web-app/cmd/pkg/models"
	"github.com/Affiction/go-web-app/cmd/pkg/render"
)

// Repos the repository used by the handlers
var Repo *Repository

// Repository represents a repository
type Repository struct {
	App *config.AppConfig
}

// Create a new repository
func NewRepository(app *config.AppConfig) *Repository {
	return &Repository{
		App: app,
	}
}

// Set repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home page handler
func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "This is template data"
	
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{StringMap: stringMap})
}

// About page handler
func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{})
}
