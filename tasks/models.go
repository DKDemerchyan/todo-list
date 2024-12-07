package tasks

import "fmt"

type Task struct {
	ID      string `json:"id,omitempty" db:"id,omitempty"`
	Date    string `json:"date"         db:"date"`
	Title   string `json:"title"        db:"title"`
	Comment string `json:"comment"      db:"comment"`
	Repeat  string `json:"repeat"       db:"repeat"`
}

func (task Task) String() string {
	return fmt.Sprintf(
		"ID: %s, Date: %s, Title: %s, Comment: %s, Repeat: %s",
		task.ID, task.Date, task.Title, task.Comment, task.Repeat)
}
