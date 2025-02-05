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
	err := s.db.QueryRow("SELECT points FROM processor WHERE id = ?", receipt_id).Scan(&points)
	if err != nil {
		log.Println("Select:", err)
		return 0, err
	}

	return points, nil
}

func (s *Store) AddPoints(points uint64) (uint64, error) {
	var result uint64
	log.Println("Points for db:", points)
	err := s.db.QueryRow("INSERT INTO processor (points) VALUES (?) RETURNING id", points).Scan(&result)
	if err != nil {
		log.Println("Insert:", err)
		return 0, err
	}
	s.CheckDB()
	return result, nil
}

func (s *Store) CheckDB() {
	rows, err := s.db.Query("SELECT * FROM processor")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var points uint64
		err := rows.Scan(&id, &points)
		log.Println(points)
		if err != nil {
			log.Fatal("Check:", err)
		}
		log.Printf("ID: %d, Points: %d\n", id, points)
	}
}
