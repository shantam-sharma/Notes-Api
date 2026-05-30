package handlers

import (
	"encoding/json"
	"net/http"
	"notes_api/internal/middleware"
	"notes_api/internal/services"
	"notes_api/internal/utils"
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
		utils.WriteError(
			w,
			http.StatusBadRequest,
			"invalid request body",
		)
		return
	}

	if req.Title == "" || req.Content == "" {
		utils.WriteError(
			w,
			http.StatusBadRequest,
			"title and content are required",
		)
		return
	}
	// 3 tasks
	// Get authenticated user ID from request context.
	// Middleware stored this after validating JWT.
	// Get authenticated user ID from request context.
	// Middleware stored this after validating JWT.
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		utils.WriteError(
			w,
			http.StatusUnauthorized,
			"unauthorized",
		)
		return
	}

	err = h.Service.CreateNote(userID, req.Title, req.Content)
	if err != nil {
		utils.WriteError(
			w,
			http.StatusInternalServerError,
			"failed to create note",
		)
		return
	}

	utils.WriteJSON(
		w,
		http.StatusCreated,
		map[string]string{
			"message": "note created",
		},
	)
}

func (h *NoteHandler) GetNotes(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)

	if !ok {
		utils.WriteError(
			w,
			http.StatusUnauthorized,
			"unauthorized",
		)
		return
	}

	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page := 1
	limit := 10

	if pageStr != "" {
		parsedPage, err := strconv.Atoi(pageStr)

		if err != nil {
			utils.WriteError(
				w,
				http.StatusBadRequest,
				"invalid page",
			)
			return
		}
		page = parsedPage
	}

	if limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)

		if err != nil {
			utils.WriteError(
				w,
				http.StatusBadRequest,
				"invalid limit",
			)
			return
		}
		limit = parsedLimit
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	if limit > 100 {
		limit = 100
	}

	notes, err := h.Service.GetNotesByUserID(
		userID,
		page,
		limit,
	)

	if err != nil {
		utils.WriteError(
			w,
			http.StatusInternalServerError,
			"failed to fetch notes",
		)
		return
	}

	utils.WriteJSON(
		w,
		http.StatusOK,
		notes,
	)
}

func (h *NoteHandler) GetNoteByID(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		utils.WriteError(
			w,
			http.StatusUnauthorized,
			"unauthorized",
		)
		return
	}

	idParam := chi.URLParam(r, "id")

	noteID, err := strconv.Atoi(idParam)

	if err != nil {
		utils.WriteError(
			w,
			http.StatusBadRequest,
			"invalid note id",
		)
		return
	}

	note, err := h.Service.GetNoteByID(noteID, userID)
	if err != nil {
		utils.WriteError(
			w,
			http.StatusNotFound,
			"note not found",
		)
		return
	}

	utils.WriteJSON(
		w,
		http.StatusOK,
		note,
	)
}

func (h *NoteHandler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	var req CreateNoteRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		utils.WriteError(
			w,
			http.StatusBadRequest,
			"invalid request body",
		)
		return
	}

	if req.Title == "" || req.Content == "" {
		utils.WriteError(
			w,
			http.StatusBadRequest,
			"title and content are required",
		)
		return
	}

	userID, ok := r.Context().Value(middleware.UserIDKey).(int)

	if !ok {
		utils.WriteError(
			w,
			http.StatusUnauthorized,
			"unauthorized",
		)
		return
	}
	//Gets the "id" value from the URL route.
	//notes/5 -> "5"
	idParam := chi.URLParam(r, "id")
	//its str like "5" so we convert to 5
	noteID, err := strconv.Atoi(idParam)

	if err != nil {
		utils.WriteError(
			w,
			http.StatusBadRequest,
			"invalid note id",
		)
		return
	}

	err = h.Service.UpdateNote(
		noteID,
		userID,
		req.Title,
		req.Content,
	)

	if err != nil {
		utils.WriteError(
			w,
			http.StatusNotFound,
			"note not found",
		)
		return
	}

	utils.WriteJSON(
		w,
		http.StatusOK,
		map[string]string{
			"message": "note updated successfully",
		},
	)
}

func (h *NoteHandler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)

	if !ok {
		utils.WriteError(
			w,
			http.StatusUnauthorized,
			"unauthorized",
		)
		return
	}

	idParam := chi.URLParam(r, "id")
	noteID, err := strconv.Atoi(idParam)
	if err != nil {
		utils.WriteError(
			w,
			http.StatusBadRequest,
			"invalid note id",
		)
		return
	}

	err = h.Service.DeleteNote(noteID, userID)
	if err != nil {
		utils.WriteError(
			w,
			http.StatusNotFound,
			"note not found",
		)
		return
	}
	utils.WriteJSON(
		w,
		http.StatusOK,
		map[string]string{
			"message": "note deleted successfully",
		},
	)
}
