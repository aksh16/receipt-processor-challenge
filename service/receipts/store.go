package receipts

import (
	"database/sql"
	"log"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetPoints(receipt_id uint64) (uint64, error) {
	var points uint64
	err := s.db.QueryRow("SELECT points FROM receipts WHERE id = ?", receipt_id).Scan(&points)
	if err != nil {
		log.Println("Select:", err)
		return 0, err
	}
	log.Printf("GetPoints ID: %d, Points: %d\n", receipt_id, points)
	return points, nil
}

func (s *Store) AddPoints(points uint64) (uint64, error) {
	var id uint64
	err := s.db.QueryRow("INSERT INTO receipts (points) VALUES (?) RETURNING id", points).Scan(&id)
	if err != nil {
		log.Println("Insert:", err)
		return 0, err
	}
	log.Printf("AddPoints ID: %d, Points: %d\n", id, points)
	return id, nil
}

func (s *Store) CheckDB() {
	rows, err := s.db.Query("SELECT * FROM receipts")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var points uint64
		err := rows.Scan(&id, &points)
		if err != nil {
			log.Fatal("Check:", err)
		}
		log.Printf("ID: %d, Points: %d\n", id, points)
	}
}
