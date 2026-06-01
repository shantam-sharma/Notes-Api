package main

import (
	"log"
	"net/http"
	"notes_api/internal/database"
	"notes_api/internal/handlers"
	"notes_api/internal/middleware"
	"notes_api/internal/repositories"
	"notes_api/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
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

	noteRepo := &repositories.NoteRepository{
		DB: db,
	}

	noteService := &services.NoteService{
		Repo: noteRepo,
	}

	noteHandler := &handlers.NoteHandler{
		Service: noteService,
	}

	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)

		r.Post("/notes", noteHandler.CreateNote)
		r.Get("/notes", noteHandler.GetNotes)
		r.Get("/notes/{id}", noteHandler.GetNoteByID)
		r.Put("/notes/{id}", noteHandler.UpdateNote)
		r.Delete("/notes/{id}", noteHandler.DeleteNote)
	})

	r.Post("/signup", authHandler.Signup)
	r.Post("/login", authHandler.Login)

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}

}
