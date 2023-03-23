package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Affiction/go-web-app/cmd/pkg/config"
	"github.com/Affiction/go-web-app/cmd/pkg/handlers"
	"github.com/Affiction/go-web-app/cmd/pkg/render"
)

const (
	port = ":8080"
)

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cant create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepository(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Server starting on port http://localhost%s\n", port)

	server := http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
