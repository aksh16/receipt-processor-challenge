package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func SQLiteStorage() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS receipts (id INTEGER PRIMARY KEY, points INTEGER)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec()
	if err != nil {
		return nil, err
	}

	if !verifyTable(db) {
		return nil, fmt.Errorf("failed to create receipts table")
	}

	return db, nil
}

func verifyTable(db *sql.DB) bool {
	var tableName string
	err := db.QueryRow(`
        SELECT name FROM sqlite_master 
        WHERE type='table' AND name='receipts'
    `).Scan(&tableName)

	if err == sql.ErrNoRows {
		return false
	}
	if err != nil {
		log.Printf("Error checking table: %v", err)
		return false
	}
	log.Println("Table created")
	return true
}
