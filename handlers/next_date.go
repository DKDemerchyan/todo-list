package handlers

import (
	"github.com/DKDemerchyan/todo-list/tasks"
	"net/http"
	"time"
)

func NextDate(writer http.ResponseWriter, r *http.Request) {
	now, _ := time.Parse("20060102", r.FormValue("now"))
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")

	nextDate, err := tasks.NextDate(now, date, repeat)
	if err != nil {
		_, _ = writer.Write([]byte(err.Error()))
	}

	_, _ = writer.Write([]byte(nextDate))
}
