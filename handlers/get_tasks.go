package handlers

import (
	"encoding/json"
	"github.com/DKDemerchyan/todo-list/database"
	"github.com/DKDemerchyan/todo-list/tasks"
	"net/http"
)

func GetTasks(ts database.TaskStore) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		allTasks, err := ts.GetAllTasks()
		if err != nil {
			http.Error(writer, errToJSON(err), http.StatusBadRequest)
		}

		if allTasks == nil {
			allTasks = []tasks.Task{}
		}

		response, err := json.Marshal(map[string][]tasks.Task{"tasks": allTasks})
		if err != nil {
			http.Error(writer, errToJSON(err), http.StatusBadRequest)
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write(response)
	}
}
