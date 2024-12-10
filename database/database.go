package database

import (
	"database/sql"
	"os"
)

func ConnectDB(dbFile string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}

	_, err = os.Stat(dbFile)
	if err != nil {
		err = CreateTableAndIdx(db)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

func CreateTableAndIdx(db *sql.DB) error {
	createQuery := `
	CREATE TABLE IF NOT EXISTS scheduler (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date TEXT NOT NULL,
		title TEXT NOT NULL,
		comment TEXT,
		repeat TEXT
	);     
	CREATE INDEX IF NOT EXISTS scheduler_date ON scheduler (date);
	`
	_, err := db.Exec(createQuery)
	if err != nil {
		return err
	}
	return nil
}
