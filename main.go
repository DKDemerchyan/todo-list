package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
)

func main() {
	router := chi.NewRouter()

	port := os.Getenv("TODO_PORT")
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}
