package routes

import (
	"github.com/a-lukianenko-dev/fem_project/internal/app"
	"github.com/go-chi/chi/v5"
)

func SetupRoutesHandler(app *app.Application) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/health", app.HandleHealthCheck)

	return router
}
