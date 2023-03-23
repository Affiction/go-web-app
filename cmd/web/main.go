package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Affiction/go-web-app/cmd/pkg/config"
	"github.com/Affiction/go-web-app/cmd/pkg/handlers"
	"github.com/Affiction/go-web-app/cmd/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const (
	port = ":8080"
)

var app config.AppConfig
var sessionManager *scs.SessionManager

func main() {

	app.Production = false

	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = app.Production

	app.Session = sessionManager

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
