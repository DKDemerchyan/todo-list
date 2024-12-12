package handlers

import (
	"github.com/DKDemerchyan/todo-list/tasks"
	"log"
	"net/http"
	"time"
)

func NextDate(writer http.ResponseWriter, r *http.Request) {
	now, err := time.Parse(tasks.DateFormat, r.FormValue("now"))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")

	nextDate, err := tasks.NextDate(now, date, repeat)
	if err != nil {
		_, _ = writer.Write([]byte(err.Error()))
	}

	_, err = writer.Write([]byte(nextDate))
	if err != nil {
		log.Println(err.Error())
	}
}
