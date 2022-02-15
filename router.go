package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kristaponis/go-mini-starter/handlers"
	"github.com/kristaponis/go-mini-starter/middlewares"
)

func router() *chi.Mux {
	r := chi.NewRouter()

	// Initialize handlers.
	static := handlers.NewStaticHandler()
	user := handlers.NewUserHandler()

	// Middleware used in all routes - global middleware.
	r.Use(middlewares.CheckUser)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Error 404 and 405 routes.
	r.NotFound(handlers.NotFound)
	r.MethodNotAllowed(handlers.MethodNotAllowed)

	// Static pages routes.
	r.Get("/", static.HomePage())
	r.Get("/contacts", static.ContactsPage)

	// User routes.
	r.Get("/user/signup", middlewares.UserLogged(user.SignupUserForm))
	r.Post("/user/signup", middlewares.UserLogged(user.SignupUser))
	r.Get("/user/login", middlewares.UserLogged(user.LoginUserForm))
	r.Post("/user/login", middlewares.UserLogged(user.LoginUser))
	r.Get("/user/dashboard", middlewares.RequireUser(user.DashboardUser))
	r.Post("/user/logout", middlewares.RequireUser(user.LogoutUser))
	r.Post("/user/delete", middlewares.RequireUser(user.DeleteUser))

	// Serve favicon icon.
	r.Get("/favicon.ico", handlers.Favicon)

	// FileServer for static files, like css, images and js.
	fs := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	return r
}
