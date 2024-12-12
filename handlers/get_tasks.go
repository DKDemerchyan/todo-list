package handlers

import (
	"encoding/json"
	"github.com/DKDemerchyan/todo-list/database"
	"github.com/DKDemerchyan/todo-list/tasks"
	"log"
	"net/http"
	"strings"
	"time"
)

const FrontendDateFormat = "02.01.2006"

func GetTasks(ts database.TaskStore) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		search := request.URL.Query().Get("search")

		// we use this variable because of disability to search by lower(RUS) in SQLite
		// and don't want to limit db response to search in this func
		var searchCase = "no search" // default
		if date, err := time.Parse(FrontendDateFormat, search); err == nil {
			searchCase = "date"
			search = date.Format(tasks.DateFormat)
		} else if search != "" {
			searchCase = "text"
		}

		allTasks, err := ts.GetTasks(search, searchCase)
		if err != nil {
			http.Error(writer, errToJSON(err), http.StatusBadRequest)
			return
		}

		if allTasks == nil {
			allTasks = []tasks.Task{}
		}

		if searchCase == "text" {
			allTasks = filterByText(allTasks, search)
		}

		response, err := json.Marshal(map[string][]tasks.Task{"tasks": allTasks})
		if err != nil {
			http.Error(writer, errToJSON(err), http.StatusBadRequest)
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

func filterByText(allTasks []tasks.Task, search string) []tasks.Task {
	var filteredTasks []tasks.Task

	for _, task := range allTasks {
		titleContainsSearch := strings.Contains(strings.ToLower(task.Title), strings.ToLower(search))
		commentContainsSearch := strings.Contains(strings.ToLower(task.Comment), strings.ToLower(search))
		if titleContainsSearch || commentContainsSearch {
			filteredTasks = append(filteredTasks, task)
		}
	}
	return filteredTasks
}
