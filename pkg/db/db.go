package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

var db *sql.DB

const schema = `
CREATE TABLE IF NOT EXISTS scheduler (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	date CHAR(8) NOT NULL DEFAULT "",
	title VARCHAR(64) NOT NULL DEFAULT "",
	comment TEXT,
	repeat VARCHAR(128) NOT NULL DEFAULT ""
	);
CREATE INDEX IF NOT EXISTS idx_scheduler_date ON scheduler(date)
`

func Init(dbFile string) error {
	install := false
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		install = true
	}

	var errOpen error
	db, errOpen = sql.Open("sqlite", dbFile)
	if errOpen != nil {
		return fmt.Errorf("Failed to open database: %v", errOpen)
	}

	if install {
		_, err := db.Exec(schema)
		if err != nil {
			_ = db.Close()
			return fmt.Errorf("Error creating table or index: %v", err)
		}
		fmt.Println("Table and index created successfully")
	}

	return nil
}

func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
