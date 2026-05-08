package main

import (
	"log"
	"notes_api/internal"
)

func main() {

	db, err := internal.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	log.Println("Server starting...")
}
