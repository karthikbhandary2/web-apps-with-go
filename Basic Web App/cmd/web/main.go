package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/karthikbhandary2/bwa/pkg/config"
	"github.com/karthikbhandary2/bwa/pkg/handlers"
	"github.com/karthikbhandary2/bwa/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	
	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}