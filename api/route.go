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

	r.HandlerFunc(http.MethodGet, "/lanjukang/", handlers.Repo.Home)

	// route for static file
	r.ServeFiles("/lanjukang/images/*filepath", http.Dir("images"))

	// user
	r.HandlerFunc(http.MethodPost, "/lanjukang/user/signup", handlers.Repo.SignupHandler)
	r.HandlerFunc(http.MethodPost, "/lanjukang/user/login", handlers.Repo.LoginHandler)
	r.HandlerFunc(http.MethodPut, "/lanjukang/user/update-role", handlers.Repo.BecomeAdminHandler)
	r.HandlerFunc(http.MethodPut, "/lanjukang/user/update-profile", handlers.Repo.UpdateUserHandler)
	r.HandlerFunc(http.MethodGet, "/lanjukang/user/detail/:id", handlers.Repo.GetUserHandler)
	r.HandlerFunc(http.MethodPost, "/lanjukang/user/forgot-password", handlers.Repo.ForgetPasswordHandler)
	r.HandlerFunc(http.MethodPost, "/lanjukang/user/reset-password", handlers.Repo.ResetPasswordHandler)
	r.HandlerFunc(http.MethodPut, "/lanjukang/user/image", handlers.Repo.UpdateUserImageHandler)

	// service
	r.HandlerFunc(http.MethodPost, "/lanjukang/service/create", handlers.Repo.CreateService)
	r.HandlerFunc(http.MethodGet, "/lanjukang/service/list", handlers.Repo.ListAllService)
	r.HandlerFunc(http.MethodGet, "/lanjukang/service/pop", handlers.Repo.ListPopService)
	r.HandlerFunc(http.MethodGet, "/lanjukang/service/detail/:id", handlers.Repo.GetServiceDetail)
	r.HandlerFunc(http.MethodPut, "/lanjukang/service/update/:id", handlers.Repo.UpdateService)
	r.HandlerFunc(http.MethodDelete, "/lanjukang/service/delete/:id", handlers.Repo.DeleteService)
	r.HandlerFunc(http.MethodGet, "/lanjukang/service/search", handlers.Repo.SeachService)
	r.HandlerFunc(http.MethodGet, "/lanjukang/service/type/:id", handlers.Repo.ServiceByType)

	// Booking
	r.HandlerFunc(http.MethodPost, "/lanjukang/booking/create", handlers.Repo.CreateBookingHandler)
	r.HandlerFunc(http.MethodPut, "/lanjukang/booking/update", handlers.Repo.UpdateBookingHandler)
	r.HandlerFunc(http.MethodGet, "/lanjukang/booking/list", handlers.Repo.GetBookingByUserHandler)
	r.HandlerFunc(http.MethodDelete, "/lanjukang/booking/delete/:id", handlers.Repo.DeleteBookingHandler)

	return middleware.EnableCors(r)
}
