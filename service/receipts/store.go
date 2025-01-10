package receipts

import (
	"database/sql"
	"fmt"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetPoints(receipt_id int64) (int64, error) {
	rows, err := s.db.Query("SELECT points FROM processor WHERE id = ?", receipt_id)
	if err != nil {
		return 0, fmt.Errorf("error querying points: %w", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return 0, fmt.Errorf("receipt not found")
	}

	var points int64
	if err := rows.Scan(&points); err != nil {
		return 0, fmt.Errorf("error scanning points: %w", err)
	}

	return points, nil
}

func (s *Store) AddReceipt(receipt string, points int) error {
	_, err := s.db.Exec("INSERT INTO processor (receipt,points) VALUES (?,?)", receipt, points)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetId() (int64, error) {
	rows, err := s.db.Query("SELECT MAX(id) FROM processor")
	if err != nil {
		return 0, fmt.Errorf("error querying id: %w", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return 0, fmt.Errorf("table is empty")
	}

	var id int64
	if err := rows.Scan(&id); err != nil {
		return 0, fmt.Errorf("error scanning id: %w", err)
	}

	return id, nil
}
