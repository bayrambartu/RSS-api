package handlers

import (
	"context"
	"database/sql"
	"net/http"
	"strings"

	"rssapi/internal/models"
)

type contextKey string

const userContextKey = contextKey("user")

func APIKeyAuthMiddleware(db *sql.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			apiKey := strings.TrimSpace(authHeader)

			if apiKey == "" {
				http.Error(w, "API key required", http.StatusUnauthorized)
				return
			}

			var user models.User
			err := db.QueryRow("SELECT id, name, email, api_key FROM users WHERE api_key = $1", apiKey).Scan(
				&user.ID, &user.Name, &user.Email, &user.APIKey,
			)

			if err != nil {
				http.Error(w, "Unvalid API key", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), userContextKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserFromContext(r *http.Request) (models.User, bool) {
	user, ok := r.Context().Value(userContextKey).(models.User)
	return user, ok
}
