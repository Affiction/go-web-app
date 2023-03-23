package config

import "html/template"

// Application configuration file
type AppConfig struct {
	TemplateCache map[string]*template.Template
}