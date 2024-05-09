package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.Handler(http.MethodGet, "/v1/movies", app.requirePermission(http.HandlerFunc(app.listMoviesHandler), "movies:read"))
	router.Handler(http.MethodPost, "/v1/movies", app.requirePermission(http.HandlerFunc(app.createMovieHandler), "movies:write"))
	router.Handler(http.MethodGet, "/v1/movies/:id", app.requirePermission(http.HandlerFunc(app.showMovieHandler), "movies:read"))
	router.Handler(http.MethodPatch, "/v1/movies/:id", app.requirePermission(http.HandlerFunc(app.updateMovieHandler), "movies:write"))
	router.Handler(http.MethodDelete, "/v1/movies/:id", app.requirePermission(http.HandlerFunc(app.deleteMovieHandler), "movies:write"))

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)

	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	return app.recoverPanic(app.enableCORS(app.rateLimit(app.authenticate(router))))
}
