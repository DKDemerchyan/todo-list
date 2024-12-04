package tasks

type Task struct {
	ID      int64  `json:"id,omitempty" db:"id,omitempty"`
	Date    string `json:"date"         db:"date"`
	Title   string `json:"title"        db:"title"`
	Comment string `json:"comment"      db:"comment"`
	Repeat  string `json:"repeat"       db:"repeat"`
}
