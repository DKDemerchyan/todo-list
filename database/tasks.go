package database

import (
	"database/sql"
	"fmt"
	"github.com/DKDemerchyan/todo-list/tasks"
)

type TaskStore struct {
	db *sql.DB
}

func NewTaskStore(db *sql.DB) TaskStore {
	return TaskStore{db: db}
}

func (ts TaskStore) CreateTask(task tasks.Task) (string, error) {
	fmt.Printf("Ща добавлю: title-%s, date-%s, comment-%s, repeat-%s\n\n", task.Title, task.Date, task.Comment, task.Repeat)
	res, err := ts.db.Exec("INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)",
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat))
	if err != nil {
		return "", err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return "", err
	}
	fmt.Println(id)
	return fmt.Sprintf("%d", id), nil
}
