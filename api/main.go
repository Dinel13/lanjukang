package main

import (
	"log"
	"net/http"

	"github.com/dinel13/lanjukang/config"
	"github.com/dinel13/lanjukang/handlers"
)

// app is a global variable for this package
var app config.AppConfig

func main() {
	app.AppName = "Lanjukang"
	app.AppVersion = "1.0.0"

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
