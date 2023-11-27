package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/SudoAaron/golang-auth0-template/platform/authenticator"
	customMiddleware "github.com/SudoAaron/golang-auth0-template/platform/middleware"
	"github.com/SudoAaron/golang-auth0-template/web/app/callback"
	"github.com/SudoAaron/golang-auth0-template/web/app/home"
	"github.com/SudoAaron/golang-auth0-template/web/app/login"
	"github.com/SudoAaron/golang-auth0-template/web/app/logout"
	"github.com/SudoAaron/golang-auth0-template/web/app/protected"
	"github.com/SudoAaron/golang-auth0-template/web/app/public"
	"github.com/SudoAaron/golang-auth0-template/web/app/user"
)

// New registers the routes and returns the router.
func New(auth *authenticator.Authenticator) *chi.Mux {
	router := chi.NewRouter()

	// Middlewares
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	authMiddleware := customMiddleware.IsAuthenticated

	// Static files
	fileServer := http.FileServer(http.Dir("web/static"))
	router.Handle("/public/*", http.StripPrefix("/public", fileServer))

	// Routes
	router.Get("/", home.Handler)
	router.Get("/login", login.Handler(auth))
	router.Get("/callback", callback.Handler(auth))
	router.Get("/logout", logout.Handler)
	router.Get("/public", public.Handler)

	// Protected routes
	router.Group(func(r chi.Router) {
		r.Use(authMiddleware)
		r.Get("/user", user.Handler)
		r.Get("/protected", protected.Handler)
	})

	return router
}
