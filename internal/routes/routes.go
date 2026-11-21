package routes

import (
	"github.com/a-lukianenko-dev/fem_project/internal/app"
	"github.com/go-chi/chi/v5"
)

func SetupRoutesHandler(app *app.Application) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/health", app.HandleHealthCheck)

	router.Get("/workouts", app.WorkoutHandler.HandleGetWorkouts)
	router.Get("/workouts/{id}", app.WorkoutHandler.HandleGetWorkoutById)
	router.Post("/workouts", app.WorkoutHandler.HandleCreateWorkout)
	router.Delete("/workouts/{id}", app.WorkoutHandler.HandleDeleteWorkoutById)

	return router
}
