package main

import (
	"net/http"

	"github.com/dinel13/lanjukang/config"
	"github.com/dinel13/lanjukang/handlers"
	"github.com/dinel13/lanjukang/middleware"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(chiMiddleware.Recoverer)
	mux.Use(middleware.TestMiddleware)

	mux.Get("/", handlers.Repo.Home)
	mux.Post("/", handlers.Repo.Home)
	mux.Post("/signup", handlers.Repo.SignupHandler)
	// mux.Post("/login", handlers.Repo.LoginHandler)

	return mux
}
