package handlers

import (
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"

	"rssapi/internal/models"

	"crypto/rand"
)

// generate 16 byte random api key
func generateAPIKey() (string, error) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func RegisterUserRoutes(r chi.Router, db *sql.DB) {
	r.Post("/users", func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "invalid data", http.StatusBadRequest)
			return
		}

		// generate apikey
		user.APIKey, err = generateAPIKey()
		if err != nil {
			http.Error(w, "API key could not be generated", http.StatusInternalServerError)
			return
		}

		// add user
		err = db.QueryRow(
			"INSERT INTO users (name, email, api_key) VALUES ($1, $2, $3) RETURNING id",
			user.Name, user.Email, user.APIKey,
		).Scan(&user.ID)

		if err != nil {
			http.Error(w, "database error", http.StatusInternalServerError)
			fmt.Println("DB Hatası:", err)
			return
		}

		fmt.Printf("Yeni kullanıcı: %+v\n", user)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	})

	// list user
	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name, email, api_key FROM users")
		if err != nil {
			http.Error(w, "database error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var users []models.User

		for rows.Next() {
			var u models.User
			if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.APIKey); err != nil {
				http.Error(w, "row scan error", http.StatusInternalServerError)
				return
			}
			users = append(users, u)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	})
}
