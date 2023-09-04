package main

import (
	"fmt"
	"net/http"
)

func handlerGetMovies(w http.ResponseWriter, r *http.Request) {
	var result []Movie
	db, err := handlerOpenDatabaseConnection()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't open database connection: %v", err))
		return
	}
	defer db.Close()

	rows, err := handlerQueryDB("SELECT ID, Title, Genre, Watched FROM movies")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't query DB: %v", err))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var movie Movie
		if err := rows.Scan(&movie.ID, &movie.Title, &movie.Genre, &movie.Watched); err != nil {
			respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't scan rows DB: %v", err))
			return
		}
		result = append(result, movie)
	}
	respondWithJSON(w, http.StatusOK, result)
}
