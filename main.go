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
	router := setupRouter(store)
	port := ":" + os.Getenv("TODO_PORT")
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}
}

func setupRouter(store database.TaskStore) http.Handler {
	webDir := "./web"
	fileServer := http.FileServer(http.Dir(webDir))

	router := chi.NewRouter()
	router.Mount("/", fileServer)

	router.Route("/api", func(r chi.Router) {
		r.Get("/nextdate", handlers.NextDate)
		r.Get("/tasks", handlers.GetTasks(store))

		r.Route("/task", func(rr chi.Router) {
			rr.Get("/", handlers.GetTask(store))
			rr.Post("/", handlers.CreateTask(store))
			rr.Put("/", handlers.UpdateTask(store))
			rr.Delete("/", handlers.DeleteTask(store))
			rr.Post("/done", handlers.TaskDone(store))
		})
	})
	return router
}
