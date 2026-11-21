package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type WorkoutHandler struct{}

func NewWorkoutHandler() *WorkoutHandler {
	return &WorkoutHandler{}
}

func (wh *WorkoutHandler) HandleGetWorkouts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "all workouts\n")
}

func (wh *WorkoutHandler) HandleGetWorkoutById(w http.ResponseWriter, r *http.Request) {
	workoutIdParam := chi.URLParam(r, "id")

	if workoutIdParam == "" {
		http.NotFound(w, r)
		return
	}

	workoutId, err := strconv.ParseInt(workoutIdParam, 10, 64)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "workout id %d\n", workoutId)
}

func (wh *WorkoutHandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "workout created\n")
}

func (wh *WorkoutHandler) HandleDeleteWorkoutById(w http.ResponseWriter, r *http.Request) {
	workoutParamId := chi.URLParam(r, "id")

	if workoutParamId == "" {
		http.NotFound(w, r)
		return
	}

	workoutId, err := strconv.ParseInt(workoutParamId, 10, 64)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Deleted workoutID: %d\n", workoutId)
}
