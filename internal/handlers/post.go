package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"rssapi/internal/models"

	"github.com/go-chi/chi/v5"
)

func RegisterPostRoutes(r chi.Router, db *sql.DB) {
	r.Group(func(protected chi.Router) {
		protected.Use(APIKeyAuthMiddleware(db))

		protected.Get("/posts", func(w http.ResponseWriter, r *http.Request) {
			// user_id yerine context’ten user çekilebilir
			user, ok := GetUserFromContext(r)
			if !ok {
				http.Error(w, "Kullanıcı bulunamadı", http.StatusInternalServerError)
				return
			}

			rows, err := db.Query(`
				SELECT posts.id, posts.title, posts.url, posts.published_at
				FROM posts
				JOIN feeds ON posts.feed_id = feeds.id
				JOIN user_feeds ON feeds.id = user_feeds.feed_id
				WHERE user_feeds.user_id = $1
				ORDER BY posts.published_at DESC
			`, user.ID)

			if err != nil {
				http.Error(w, "databse error", http.StatusInternalServerError)
				fmt.Println("POSTS query error\n:", err)
				return
			}
			defer rows.Close()

			var posts []models.Post

			for rows.Next() {
				var p models.Post
				if err := rows.Scan(&p.ID, &p.Title, &p.URL, &p.PublishedAt); err != nil {
					http.Error(w, "Post could not be readed", http.StatusInternalServerError)
					return
				}
				posts = append(posts, p)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(posts)
		})
	})
}
