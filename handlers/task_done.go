package handlers

import (
	"github.com/DKDemerchyan/todo-list/database"
	"github.com/DKDemerchyan/todo-list/tasks"
	"net/http"
	"time"
)

func TaskDone(ts database.TaskStore) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		id := request.URL.Query().Get("id")
		task, err := ts.GetTaskByID(id)
		if err != nil {
			http.Error(writer, errToJSON(err), http.StatusBadRequest)
			return
		}

		if task.Repeat == "" {
			err = ts.DeleteTask(task.ID)
			if err != nil {
				http.Error(writer, errToJSON(err), http.StatusBadRequest)
				return
			}
		} else {
			nextDate, err := tasks.NextDate(time.Now(), task.Date, task.Repeat)
			if err != nil {
				http.Error(writer, errToJSON(err), http.StatusBadRequest)
				return
			}
			task.Date = nextDate
			err = ts.UpdateTask(task)
			if err != nil {
				http.Error(writer, errToJSON(err), http.StatusBadRequest)
				return
			}
		}

		writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write([]byte("{}"))
	}
}
