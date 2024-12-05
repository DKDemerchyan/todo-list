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

	return fmt.Sprintf("%d", id), nil
}

func (ts TaskStore) GetAllTasks() ([]tasks.Task, error) {
	rows, err := ts.db.Query("SELECT * FROM scheduler ORDER BY date ASC LIMIT 50")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []tasks.Task
	for rows.Next() {
		t := tasks.Task{}
		err = rows.Scan(&t.ID, &t.Date, &t.Title, &t.Comment, &t.Repeat)
		if err != nil {
			return nil, err
		}
		res = append(res, t)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return res, nil
}
