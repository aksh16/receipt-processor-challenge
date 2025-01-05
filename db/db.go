package db

import (
	"database/sql"
	"log"
)

func SQLiteStorage() (*sql.DB, error) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
