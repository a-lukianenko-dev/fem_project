package store

import (
	"database/sql"
	"fmt"
	"io/fs"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("pgx", "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")

	if err != nil {
		return nil, fmt.Errorf("db: open %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("db: open %w", err)
	}

	fmt.Println("Connected to the database.")
	return db, nil
}

func Migrate(embedMigrations fs.FS, db *sql.DB) error {
	goose.SetBaseFS(embedMigrations)

	// reset the FS after migration
	defer func() {
		goose.SetBaseFS(nil)
	}()

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("goose set dialect: %w", err)
	}

	if err := goose.Up(db, "."); err != nil {
		return fmt.Errorf("goose up: %w", err)
	}

	return nil
}
