package handlers

import (
	"encoding/json"
	"net/http"

	"notes_api/internal/services"
	"notes_api/internal/utils"
)

// incoming json structure
type SignupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthHandler struct {
	AuthService *services.AuthService
}

// (h *AuthHandler) -> method receiver (this fun belong to auth handler)
func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	// after using chi this is not necesseraly required anymore but we are keeping it for now
	if r.Method != http.MethodPost {
		utils.WriteError(
			w,
			http.StatusMethodNotAllowed,
			"method not allowed",
		)
		return
	}
	// body can contain large data so good practice to close after use
	defer r.Body.Close()

	var req SignupRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	// Basic validation
	if req.Email == "" || req.Name == "" || req.Password == "" {
		utils.WriteError(w, http.StatusBadRequest, "All fields are required")
		return
	}

	err = h.AuthService.Signup(
		req.Name,
		req.Email,
		req.Password,
	)

	if err != nil {
		if err.Error() == "user already exists" {
			utils.WriteError(
				w,
				http.StatusConflict,
				err.Error(),
			)
			return
		}

		utils.WriteError(
			w,
			http.StatusInternalServerError,
			"failed to create user",
		)
		return
	}
	utils.WriteJSON(
		w,
		http.StatusCreated,
		map[string]string{
			"message": "user created successfully",
		},
	)

}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteError(
			w,
			http.StatusMethodNotAllowed,
			"method not allowed",
		)
		return
	}

	defer r.Body.Close()

	var req LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Email == "" || req.Password == "" {
		utils.WriteError(w, http.StatusBadRequest, "All fields are required")
		return
	}

	token, err := h.AuthService.Login(
		req.Email,
		req.Password,
	)

	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	utils.WriteJSON(
		w,
		http.StatusOK,
		map[string]string{
			"token": token,
		},
	)
}
