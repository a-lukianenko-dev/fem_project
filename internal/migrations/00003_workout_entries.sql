-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS workout_entries (
  id BIGSERIAL PRIMARY KEY,
  workout_id BIGINT NOT NULL REFERENCES workouts(id) ON DELETE CASCADE,
  exercise_name VARCHAR(255) NOT NULL,
  sets INTEGER NOT NULL,
  reps INTEGER,
  duration_seconds INTEGER,
  weight DECIMAL(5, 2),
  notes TEXT,
  order_index INTEGER NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT workout_entry_exactly_one_metric CHECK (
    num_nonnulls(reps, duration_seconds) = 1
  ),
  CONSTRAINT workout_entry_weight_non_negative CHECK (
    weight IS NULL OR weight >= 0
  ),
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE workout_entries;
-- +goose StatementEnd