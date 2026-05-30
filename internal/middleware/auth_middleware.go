package middleware

import (
	"context"
	"net/http"
	"notes_api/internal/utils"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// Custom context key type
type contextKey string

// Context key constant
const UserIDKey contextKey = "user_id"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// STEP 1 — Get Authorization Header
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			utils.WriteError(
				w,
				http.StatusUnauthorized,
				"missing authorization header",
			)
			return
		}

		// STEP 2 — Validate Bearer Format
		parts := strings.SplitN(authHeader, " ", 2)

		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.WriteError(
				w,
				http.StatusUnauthorized,
				"invalid authorization format",
			)
			return
		}

		// STEP 3 — Extract Token
		tokenString := parts[1]

		// STEP 4 — Parse JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			// Validate signing method
			_, ok := token.Method.(*jwt.SigningMethodHMAC)

			if !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			// Load JWT secret safely
			jwtSecret := os.Getenv("JWT_SECRET")

			if jwtSecret == "" {
				return nil, jwt.ErrSignatureInvalid
			}

			return []byte(jwtSecret), nil
		})

		// STEP 5 — Validate Token
		if err != nil || !token.Valid {
			utils.WriteError(
				w,
				http.StatusUnauthorized,
				"invalid token",
			)
			return
		}

		// STEP 6 — Extract Claims
		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			utils.WriteError(
				w,
				http.StatusUnauthorized,
				"invalid token claims",
			)
			return
		}

		// STEP 7 — Extract user_id safely
		userIDClaim, exists := claims["user_id"]

		if !exists {
			utils.WriteError(
				w,
				http.StatusUnauthorized,
				"missing user_id",
			)
			return
		}

		userIDFloat, ok := userIDClaim.(float64)

		if !ok {
			utils.WriteError(
				w,
				http.StatusUnauthorized,
				"invalid user_id",
			)
			return
		}

		userID := int(userIDFloat)

		// STEP 8 — Store user_id in context
		ctx := context.WithValue(r.Context(), UserIDKey, userID)

		// Attach updated context to request
		r = r.WithContext(ctx)

		// STEP 9 — Continue Request
		next.ServeHTTP(w, r)
	})
}
