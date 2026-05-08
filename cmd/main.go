package main

import (
	"log"
	"notes_api/internal/database"
)

func main() {

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	log.Println("Server starting...")
}
