package store

import "database/sql"

/* Store: abstraction over the database
   store can:
	 - createWorkout
	 - getWorkoutById
	 - getWorkouts
	 - updateWorkout
	 - DeleteWorkout

	 Data types:
	 - Workout
	 - WorkoutEntry
	 - WorkoutStore + constructor
*/

type Workout struct {
	Id              int            `json:"id"`
	Title           string         `json:"title"`
	Description     *string        `json:"description"`
	WorkoutEntries  []WorkoutEntry `json:"workout_entries"`
	DurationMinutes int            `json:"duration_minutes"`
	CaloriesBurned  *int           `json:"calories_burned"`
}

type WorkoutEntry struct {
	Id              int      `json:"id"`
	ExerciseName    string   `json:"exercise_name"`
	Sets            int      `json:"sets"`
	Reps            *int     `json:"reps"`
	DurationSeconds *int     `json:"duration_seconds"`
	Weight          *float64 `json:"weight"`
	Notes           *string  `json:"notes"`
	OrderIndex      int      `json:"order_index"`
}

type WorkoutStore interface {
	CreateWorkout(workout *Workout) (*Workout, error)
	GetWorkoutById(id int) (*Workout, error)
	// GetWorkouts() (*[]Workout, error)
	// UpdateWorkout(id int) (*Workout, error)
	// DeleteWorkout(id int) (*Workout, error)
}

type PostgresWorkoutStore struct {
	DB *sql.DB
}

func NewWorkoutStore(db *sql.DB) WorkoutStore {
	return &PostgresWorkoutStore{
		DB: db,
	}
}

func (pgws *PostgresWorkoutStore) CreateWorkout(workout *Workout) (*Workout, error) {
	tx, err := pgws.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	workoutQuery := `
		INSERT INTO workouts (title, description, duration_minutes, calories_burned)
		VALUES($1, $2, $3, $4)
		RETURNING id;
	`
	err = tx.QueryRow(workoutQuery, workout.Title, workout.Description, workout.DurationMinutes, workout.CaloriesBurned).Scan(&workout.Id)
	if err != nil {
		return nil, err
	}

	workoutEntryQuery := `
		INSERT INTO workout_entries (workout_id, exercise_name, sets, reps, duration_seconds, weight, notes, order_index)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id;
	`
	for i, entry := range workout.WorkoutEntries {
		var workoutEntryId int
		err = tx.QueryRow(workoutEntryQuery, workout.Id, entry.ExerciseName, entry.Sets, entry.Reps, entry.DurationSeconds, entry.Weight, entry.Notes, entry.OrderIndex).Scan(&workoutEntryId)

		if err != nil {
			return nil, err
		}
		workout.WorkoutEntries[i].Id = workoutEntryId
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return workout, nil
}

func (pgws *PostgresWorkoutStore) GetWorkoutById(id int) (*Workout, error) {
	return &Workout{}, nil
}
