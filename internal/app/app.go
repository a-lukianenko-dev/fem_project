package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/a-lukianenko-dev/fem_project/internal/api"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	workoutHandler := api.NewWorkoutHandler()

	application := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
	}

	return application, nil
}

func (app *Application) HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is available\n")
}
