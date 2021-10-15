package main

import (
	"net/http"

	"github.com/dinel13/lanjukang/config"
	"github.com/dinel13/lanjukang/handlers"
	"github.com/dinel13/lanjukang/middleware"
	"github.com/julienschmidt/httprouter"
)

func routes(app *config.AppConfig) http.Handler {
	r := httprouter.New()

	// make a recover midleware
	// mux.Use(chiMiddleware.Recoverer)

	r.HandlerFunc(http.MethodGet, "/", handlers.Repo.Home)

	// user
	r.HandlerFunc(http.MethodPost, "/signup", handlers.Repo.SignupHandler)
	r.HandlerFunc(http.MethodPost, "/login", handlers.Repo.LoginHandler)
	r.HandlerFunc(http.MethodPost, "/update-role", handlers.Repo.BecomeAdminHandler)

	// service
	r.HandlerFunc(http.MethodPost, "/service/create", handlers.Repo.CreateService)
	r.HandlerFunc(http.MethodGet, "/service/list", handlers.Repo.ListAllService)
	r.HandlerFunc(http.MethodGet, "/service/detail/:id", handlers.Repo.GetServiceDetail)
	r.HandlerFunc(http.MethodPut, "/service/update/:id", handlers.Repo.UpdateService)
	r.HandlerFunc(http.MethodDelete, "/service/delete/:id", handlers.Repo.DeleteService)

	return middleware.EnableCors(r)
}
