package main

import (
	"backend/cmd/api"
	"backend/db"
	"log"
)

func main() {
	db, err := db.SQLiteStorage()
	if err != nil {
		log.Fatal(err)
	}
	server := api.NewAPIServer(":8080", nil)
	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}

}
