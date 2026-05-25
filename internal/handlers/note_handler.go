package handlers

import (
	"encoding/json"
	"net/http"
	"notes_api/internal/middleware"
	"notes_api/internal/services"
)

type NoteHandler struct {
	Service *services.NoteService
}
type CreateNoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (h *NoteHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var req CreateNoteRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.Title == "" || req.Content == "" {
		http.Error(w, "title and content are required", http.StatusBadRequest)
		return
	}
	// 3 tasks
	// Get authenticated user ID from request context.
	// Middleware stored this after validating JWT.
	// Get authenticated user ID from request context.
	// Middleware stored this after validating JWT.
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	err = h.Service.CreateNote(userID, req.Title, req.Content)
	if err != nil {
		http.Error(w, "failed to create note", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "note created",
	})
}
