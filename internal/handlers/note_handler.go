package handlers

import (
	"encoding/json"
	"net/http"
	"notes_api/internal/middleware"
	"notes_api/internal/services"
	"strconv"

	"github.com/go-chi/chi/v5"
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

func (h *NoteHandler) GetNotes(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)

	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	notes, err := h.Service.GetNotesByUserID(userID)

	if err != nil {
		http.Error(w, "failed to fetch notes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(notes)
}

func (h *NoteHandler) GetNoteByID(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	idParam := chi.URLParam(r, "id")

	noteID, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "invalid note id", http.StatusBadRequest)
		return
	}

	note, err := h.Service.GetNoteByID(noteID, userID)
	if err != nil {
		http.Error(w, "note not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func (h *NoteHandler) UpdateNote(w http.ResponseWriter, r *http.Request) {
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

	userID, ok := r.Context().Value(middleware.UserIDKey).(int)

	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	//Gets the "id" value from the URL route.
	//notes/5 -> "5"
	idParam := chi.URLParam(r, "id")
	//its str like "5" so we convert to 5
	noteID, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "invalid note id", http.StatusBadRequest)
		return
	}

	err = h.Service.UpdateNote(
		noteID,
		userID,
		req.Title,
		req.Content,
	)

	if err != nil {
		http.Error(w, "note not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"message": "note updated successfully",
	})
}

func (h *NoteHandler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)

	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	idParam := chi.URLParam(r, "id")
	noteID, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "invalid note id", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteNote(noteID, userID)
	if err != nil {
		http.Error(w, "note not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"message": "note deleted successfully",
	})
}
