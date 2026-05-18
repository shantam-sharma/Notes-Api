package main

import (
	"log"
	"net/http"
	"notes_api/internal/database"
	"notes_api/internal/handlers"
	"notes_api/internal/repositories"
	"notes_api/internal/services"

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

	userRepo := repositories.UserRepository{
		DB: db,
	}

	authService := services.AuthService{
		UserRepo: &userRepo,
	}

	authHandler := handlers.AuthHandler{
		AuthService: &authService,
	}

	http.HandleFunc("/signup", authHandler.Signup)
	http.HandleFunc("/login", authHandler.Login)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
