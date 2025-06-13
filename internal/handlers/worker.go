package handlers

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
)

func StartFeedWorker(db *sql.DB, interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		for range ticker.C {
			fmt.Println("Worker triggered, feed's scanning...")

			rows, err := db.Query("SELECT id, url FROM feeds")
			if err != nil {
				fmt.Println("Feed query error:", err)
				continue
			}
			defer rows.Close()

			for rows.Next() {
				var feedID int
				var feedURL string
				if err := rows.Scan(&feedID, &feedURL); err != nil {
					fmt.Println("line could not be readed :", err)
					continue
				}

				fp := gofeed.NewParser()
				parsedFeed, err := fp.ParseURL(feedURL)
				if err != nil {
					fmt.Printf("Feed could not be received (%s): %v\n", feedURL, err)
					continue
				}

				for _, item := range parsedFeed.Items {
					_, err := db.Exec(
						"INSERT INTO posts (feed_id, title, url, published_at) VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING",
						feedID,
						item.Title,
						item.Link,
						item.PublishedParsed,
					)
					if err != nil {
						fmt.Println("Post could not be added:", err)
					}
				}
			}

			fmt.Println("Worker completed")
		}
	}()
}
