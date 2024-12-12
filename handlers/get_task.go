package handlers

import (
	"encoding/json"
	"errors"
	"github.com/DKDemerchyan/todo-list/database"
	"log"
	"net/http"
)

const clientGetTaskErr = "ошибка при получении задачи"

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
			log.Printf(err.Error())
			err := errors.New(clientGetTaskErr)
			http.Error(writer, errToJSON(err), http.StatusInternalServerError)
			return
		}

		response, err := json.Marshal(task)
		if err != nil {
			log.Printf(err.Error())
			err := errors.New(clientGetTaskErr)
			http.Error(writer, errToJSON(err), http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		writer.WriteHeader(http.StatusOK)
		_, err = writer.Write(response)
		if err != nil {
			log.Printf(err.Error())
		}
	}
}
