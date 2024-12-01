package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/DKDemerchyan/todo-list/database"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	appPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dbFile := filepath.Join(appPath, os.Getenv("TODO_DBFILE"))

	db, err := database.ConnectDB(dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	webDir := "./web"

	mux.Handle("/", http.FileServer(http.Dir(webDir)))

	port := ":" + os.Getenv("TODO_PORT")
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}
