package handlers

import (
	"github.com/DKDemerchyan/todo-list/database"
	"net/http"
)

func DeleteTask(ts database.TaskStore) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		id := request.URL.Query().Get("id")
		if err := ts.DeleteTask(id); err != nil {
			http.Error(writer, errToJSON(err), http.StatusBadRequest)
			return
		}

		writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write([]byte("{}"))
	}
}
