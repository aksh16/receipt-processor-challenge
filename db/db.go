package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func SQLiteStorage() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory")
	if err != nil {
		log.Fatal(err)
	}
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS processor (id INTEGER PRIMARY KEY, receipt TEXT, points INTEGER)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	return db, nil
}
