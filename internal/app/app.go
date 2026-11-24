package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/a-lukianenko-dev/fem_project/internal/api"
	"github.com/a-lukianenko-dev/fem_project/internal/migrations"
	"github.com/a-lukianenko-dev/fem_project/internal/store"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB             *sql.DB
}

func NewApplication() (*Application, error) {
	db, err := store.Open()
	if err != nil {
		return nil, err
	}

	err = store.Migrate(migrations.FS, db)
	if err != nil {
		return nil, err
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	workoutHandler := api.NewWorkoutHandler()

	application := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
		DB:             db,
	}

	return application, nil
}

func (app *Application) HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is available\n")
}
