package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/DKDemerchyan/todo-list/database"
	"github.com/DKDemerchyan/todo-list/tasks"
	"net/http"
	"time"
)

type TaskID struct {
	ID string `json:"id"`
}

func CreateTask(ts database.TaskStore) http.HandlerFunc {
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

		if err = validateTask(&task); err != nil {
			http.Error(writer, errToJSON(err), http.StatusBadRequest)
			return
		}

		id, err := ts.CreateTask(task)
		if err != nil {
			http.Error(writer, errToJSON(err), http.StatusBadRequest)
			return
		}

		response, err := json.Marshal(TaskID{ID: id})
		if err != nil {
			http.Error(writer, errToJSON(err), http.StatusBadRequest)
			return
		}

		writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		writer.WriteHeader(http.StatusCreated)
		_, _ = writer.Write(response)
	}
}

func validateTask(task *tasks.Task) error {
	if len(task.Title) < 1 {
		return errors.New("title must not be empty")
	}

	dateToday := time.Now().Format(tasks.DateFormat)
	if task.Date == "" {
		task.Date = dateToday
	}

	if len(task.Date) != 8 {
		return errors.New("date format is wrong: less than 8 characters")
	}

	taskDate, err := time.Parse(tasks.DateFormat, task.Date)
	if err != nil {
		return errors.New("date is not correct")
	}

	if taskDate.Before(time.Now()) && task.Repeat == "" {
		task.Date = dateToday
	}

	if len(task.Repeat) > 0 {
		nextDate, err := tasks.NextDate(time.Now(), task.Date, task.Repeat)
		if err != nil {
			return err
		}

		if taskDate.Format(tasks.DateFormat) < time.Now().Format(tasks.DateFormat) {
			task.Date = nextDate
		}
	}

	return nil
}

func errToJSON(err error) string {
	jsonErr, err := json.Marshal(map[string]string{"error": err.Error()})
	if err != nil {
		return ""
	}
	return string(jsonErr)
}
