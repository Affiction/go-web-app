package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

// Application configuration file
type AppConfig struct {
	Production    bool
	UseCache      bool
	TemplateCache map[string]*template.Template
	Session       *scs.Session
}
