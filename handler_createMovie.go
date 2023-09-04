package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handlerCreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid JSON: %v", err))
		return
	}

	db, err := handlerOpenDatabaseConnection()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to connect to database: %v", err))
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO movies(id, title, genre, watched) VALUES($1, $2, $3, $4)")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't prepare statement: %v", err))
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(movie.ID, movie.Title, movie.Genre, false)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't execute statement: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, movie)
}
