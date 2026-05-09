package main

import (
	"log"
	"notes_api/internal/database"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("err loading .env file")
	}
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	// defer is processed first but not executed until the app closes.
	defer db.Close()

	log.Println("Server starting...")
}
