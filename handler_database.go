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

func handlerQueryDB(query string) (rows *sql.Rows, err error) {
	db, err := handlerOpenDatabaseConnection()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err = db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return rows, nil
}
