package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	mux := http.NewServeMux()

	webDir := "./web"

	mux.Handle("/", http.FileServer(http.Dir(webDir)))

	err := godotenv.Load()
	port := ":" + os.Getenv("TODO_PORT")
	if err = http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}
