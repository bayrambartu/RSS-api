package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"rssapi/internal/models"

	"github.com/go-chi/chi/v5"
	"github.com/mmcdole/gofeed"
)

func RegisterFeedRoutes(r chi.Router, db *sql.DB) {
	r.Group(func(protected chi.Router) {
		protected.Use(APIKeyAuthMiddleware(db))

		// add Feed
		protected.Post("/feeds", func(w http.ResponseWriter, r *http.Request) {
			var input struct {
				URL string `json:"url"`
			}

			err := json.NewDecoder(r.Body).Decode(&input)
			if err != nil {
				http.Error(w, "invalid JSON", http.StatusBadRequest)
				return
			}

			fp := gofeed.NewParser()
			parsedFeed, err := fp.ParseURL(input.URL)
			if err != nil {
				http.Error(w, "RSS could not be pulled", http.StatusBadRequest)
				fmt.Println("RSS error:", err)
				return
			}

			var feed models.Feed
			err = db.QueryRow(
				"INSERT INTO feeds (url, title) VALUES ($1, $2) RETURNING id, url, title",
				input.URL,
				parsedFeed.Title,
			).Scan(&feed.ID, &feed.URL, &feed.Title)

			if err != nil {
				http.Error(w, "database error", http.StatusInternalServerError)
				fmt.Println("DB INSERT error:", err)
				return
			}

			fmt.Printf("new feed added: %+v\n", feed)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(feed)
		})

		protected.Post("/subscriptions", func(w http.ResponseWriter, r *http.Request) {
			var input struct {
				UserID int `json:"user_id"`
				FeedID int `json:"feed_id"`
			}

			err := json.NewDecoder(r.Body).Decode(&input)
			if err != nil {
				http.Error(w, "invalid JSON", http.StatusBadRequest)
				return
			}

			_, err = db.Exec(
				"INSERT INTO user_feeds (user_id, feed_id) VALUES ($1, $2) ON CONFLICT DO NOTHING",
				input.UserID,
				input.FeedID,
			)

			if err != nil {
				http.Error(w, "database error", http.StatusInternalServerError)
				fmt.Println("Subscription error:", err)
				return
			}

			fmt.Printf("user %d → Feed %d subscription could be saved\n", input.UserID, input.FeedID)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "subscriptions successful",
			})
		})
	})
}
