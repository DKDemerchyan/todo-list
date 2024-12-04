package handlers

import (
	"github.com/DKDemerchyan/todo-list/tasks"
	"net/http"
	"time"
)

func NextDate(writer http.ResponseWriter, r *http.Request) {
	strNow := r.FormValue("now")
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")

	now, err := time.Parse(tasks.DateFormat, strNow)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	nextDate, err := tasks.NextDate(now, date, repeat)
	if err != nil {
		_, _ = writer.Write([]byte(err.Error()))
	}

	_, _ = writer.Write([]byte(nextDate))
}
