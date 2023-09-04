package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type Movie struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Watched bool   `json:"watched"`
	Genre   string `json:"genre"`
}

var movies []Movie

func main() {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	movies = append(movies, Movie{
		ID:      "1",
		Title:   "Movie example",
		Watched: false,
		Genre:   "Action",
	})

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerHealthz)
	v1Router.Get("/movies", handlerGetMovies)
	v1Router.Post("/movies", handlerCreateMovie)

	router.Mount("/v1", v1Router)

	fmt.Printf("Starting server at port 8000")
	srv := &http.Server{
		Handler: router,
		Addr:    ":8000",
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
