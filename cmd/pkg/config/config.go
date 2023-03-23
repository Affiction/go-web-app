package config

import "html/template"

// Application configuration file
type AppConfig struct {
	UseCache bool
	TemplateCache map[string]*template.Template
}