package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func handlerDeleteMovie(w http.ResponseWriter, r *http.Request) {
	movieID := chi.URLParam(r, "movieID")
	db, err := handlerOpenDatabaseConnection()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to connect to database: %v", err))
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM movies WHERE ID = ?")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to DELETE movie: %v", err))
	}
	defer stmt.Close()

	_, err = stmt.Exec(movieID)
	if err != nil {
	}

}
