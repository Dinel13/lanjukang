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

	// route for static file
	r.ServeFiles("/images/*filepath", http.Dir("images"))

	// user
	r.HandlerFunc(http.MethodPost, "/user/signup", handlers.Repo.SignupHandler)
	r.HandlerFunc(http.MethodPost, "/user/login", handlers.Repo.LoginHandler)
	r.HandlerFunc(http.MethodPut, "/user/update-role", handlers.Repo.BecomeAdminHandler)
	r.HandlerFunc(http.MethodPut, "/user/update-profile", handlers.Repo.UpdateUserHandler)
	r.HandlerFunc(http.MethodGet, "/user/detail/:id", handlers.Repo.GetUserHandler)
	r.HandlerFunc(http.MethodPost, "/user/forgot-password", handlers.Repo.ForgetPasswordHandler)
	r.HandlerFunc(http.MethodPost, "/user/reset-password", handlers.Repo.ResetPasswordHandler)
	r.HandlerFunc(http.MethodPut, "/user/image", handlers.Repo.UpdateUserImageHandler)

	// service
	r.HandlerFunc(http.MethodPost, "/service/create", handlers.Repo.CreateService)
	r.HandlerFunc(http.MethodGet, "/service/list", handlers.Repo.ListAllService)
	r.HandlerFunc(http.MethodGet, "/service/pop", handlers.Repo.ListPopService)
	r.HandlerFunc(http.MethodGet, "/service/detail/:id", handlers.Repo.GetServiceDetail)
	r.HandlerFunc(http.MethodPut, "/service/update/:id", handlers.Repo.UpdateService)
	r.HandlerFunc(http.MethodDelete, "/service/delete/:id", handlers.Repo.DeleteService)
	r.HandlerFunc(http.MethodGet, "/service/search", handlers.Repo.SeachService)

	// Booking
	r.HandlerFunc(http.MethodPost, "/booking/create", handlers.Repo.CreateBookingHandler)
	r.HandlerFunc(http.MethodPut, "/booking/update", handlers.Repo.UpdateBookingHandler)
	r.HandlerFunc(http.MethodGet, "/booking/list", handlers.Repo.GetBookingByUserHandler)
	r.HandlerFunc(http.MethodDelete, "/booking/delete/:id", handlers.Repo.DeleteBookingHandler)

	return middleware.EnableCors(r)
}
