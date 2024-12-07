package handlers

import (
	"encoding/json"
	"errors"
	"github.com/DKDemerchyan/todo-list/database"
	"net/http"
)

func GetTask(ts database.TaskStore) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		id := request.URL.Query().Get("id")
		if id == "" {
			err := errors.New("task id is required")
			http.Error(writer, errToJSON(err), http.StatusBadRequest)
			return
		}

		task, err := ts.GetTaskByID(id)
		if err != nil {
			http.Error(writer, errToJSON(err), http.StatusBadRequest)
			return
		}

		response, err := json.Marshal(task)
		if err != nil {
			http.Error(writer, errToJSON(err), http.StatusBadRequest)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write(response)
	}
}
