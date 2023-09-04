package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func handlerOpenDatabaseConnection() (*sql.DB, error) {
	dbURL := handlerGetEnv()

	db, err := sql.Open("sqlite3", dbURL)
	if err != nil {
		return nil, err
	}

	return db, nil
}
