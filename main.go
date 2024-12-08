package main

import (
	"github.com/DKDemerchyan/todo-list/database"
	"github.com/DKDemerchyan/todo-list/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Database connection
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
	store := database.NewTaskStore(db)

	// Routing
	webDir := "./web"
	fileServer := http.FileServer(http.Dir(webDir))

	router := chi.NewRouter()

	router.Mount("/", fileServer)
	router.Get("/api/nextdate", handlers.NextDate)
	router.Get("/api/task", handlers.GetTask(store))
	router.Post("/api/task", handlers.CreateTask(store))
	router.Put("/api/task", handlers.UpdateTask(store))
	router.Delete("/api/task", handlers.DeleteTask(store))
	router.Post("/api/task/done", handlers.TaskDone(store))
	router.Get("/api/tasks", handlers.GetTasks(store))

	port := ":" + os.Getenv("TODO_PORT")
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}
}
