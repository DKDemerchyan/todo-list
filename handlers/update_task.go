package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/DKDemerchyan/todo-list/database"
	"github.com/DKDemerchyan/todo-list/tasks"
	"net/http"
)

func UpdateTask(ts database.TaskStore) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var task tasks.Task
		var buf bytes.Buffer

		_, err := buf.ReadFrom(request.Body)
		if err != nil {
			http.Error(writer, errToJSON(err), http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(buf.Bytes(), &task)
		if err != nil {
			http.Error(writer, errToJSON(err), http.StatusBadRequest)
			return
		}

		if err = validateTaskOnCreate(&task); err != nil {
			http.Error(writer, errToJSON(err), http.StatusBadRequest)
			return
		}

		if _, err = ts.GetTaskByID(task.ID); err != nil {
			http.Error(writer, errToJSON(err), http.StatusBadRequest)
			return
		}

		err = ts.UpdateTask(task)
		if err != nil {
			http.Error(writer, errToJSON(err), http.StatusBadRequest)
			return
		}

		writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write([]byte("{}"))
	}
}
