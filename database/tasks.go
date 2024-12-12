package database

import (
	"database/sql"
	"errors"
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

func (ts TaskStore) GetTasks(search string, searchCase string) ([]tasks.Task, error) {
	var resLimit = 50
	var rows *sql.Rows
	var err error

	switch searchCase { // we return all tasks for text search cases because of disability to LOWER(RUS)
	case "no search":
		rows, err = ts.db.Query("SELECT * FROM scheduler ORDER BY date ASC LIMIT ?", resLimit)
	case "text":
		rows, err = ts.db.Query("SELECT * FROM scheduler ORDER BY date ASC")
	case "date":
		rows, err = ts.db.Query("SELECT * FROM scheduler WHERE date LIKE ? ORDER BY date ASC LIMIT ?",
			search, resLimit)
	}

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

func (ts TaskStore) GetTaskByID(id string) (tasks.Task, error) {
	row := ts.db.QueryRow("SELECT * FROM scheduler WHERE id = $1", id)

	task := tasks.Task{}
	err := row.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return task, errors.New("there is no task with that id")
	}

	return task, nil
}

func (ts TaskStore) UpdateTask(task tasks.Task) error {
	_, err := ts.db.Exec(
		"UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat WHERE id = :id",
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
		sql.Named("id", task.ID))
	if err != nil {
		return err
	}
	return nil
}

func (ts TaskStore) DeleteTask(id string) error {
	res, err := ts.db.Exec("DELETE FROM scheduler WHERE id = $1", id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("there is no task with that id")
	}

	return nil
}
