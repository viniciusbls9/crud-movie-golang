package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func handlerUpdateMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid JSON: %v", err), http.StatusBadRequest)
		return
	}

	movieID := chi.URLParam(r, "movieID")

	db, err := handlerOpenDatabaseConnection()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to connect to database: %v", err))
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE movies SET TITLE = ?, GENRE = ?, WATCHED = ? WHERE ID = ?")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't prepare statement: %v", err))
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(movie.Title, movie.Genre, movie.Watched, movieID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't execute statement: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, movie)
}
