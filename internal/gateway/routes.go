package gateway

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Handlers struct {
	HealthCheckHandler  http.HandlerFunc
	ShowMovieHandler    http.HandlerFunc
	CreateMovieHandler  http.HandlerFunc
	UpdateMovieHandler  http.HandlerFunc
	DeleteMovieHandler  http.HandlerFunc
	ListMoviesHandler   http.HandlerFunc
	RegisterUserHandler http.HandlerFunc
	ActivateUserHandler http.HandlerFunc
}

func Routes(handlers Handlers) http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", handlers.HealthCheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", handlers.ShowMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies", handlers.ListMoviesHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", handlers.CreateMovieHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/movies/:id", handlers.UpdateMovieHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/movies/:id", handlers.DeleteMovieHandler)

	router.HandlerFunc(http.MethodPost, "/v1/users", handlers.RegisterUserHandler)

	router.HandlerFunc(http.MethodPut, "/v1/users/activated", handlers.ActivateUserHandler)

	return recoverPanic(rateLimit(router))
}
