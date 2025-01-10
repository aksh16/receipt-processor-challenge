package db

import (
	"database/sql"
	"log"
)

func SQLiteStorage() (*sql.DB, error) {
	db, err := sql.Open("sqlite", ":memory")
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS processor (id INTEGER PRIMARY KEY, receipt TEXT, points INTEGER)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	return db, nil
}
