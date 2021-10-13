package main

import (
	"log"
	"net/http"

	"github.com/dinel13/lanjukang/config"
	"github.com/dinel13/lanjukang/handlers"
	"github.com/dinel13/lanjukang/middleware"
	"github.com/go-chi/chi"
)

// app is a global variable for this package
var app config.AppConfig

func main() {

	app.AppName = "Lanjukang"
	app.AppVersion = "1.0.0"
	app.AppPort = "8080"

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	mux := chi.NewRouter()

	mux.Use(middleware.TestMiddleware)

	mux.Get("/", handlers.Repo.Home)
	mux.Post("/users", handlers.Repo.SignupHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
