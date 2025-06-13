package main

import (
	"fmt"
	"github.com/go-chi/cors"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"rssapi/internal/db"
	"rssapi/internal/handlers"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer database.Close()

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3001"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	}))

	handlers.RegisterUserRoutes(r, database)
	handlers.RegisterFeedRoutes(r, database)
	handlers.RegisterPostRoutes(r, database)

	handlers.StartFeedWorker(database, 30*time.Second)

	fmt.Println(" Server is running: http://localhost:3000")
	http.ListenAndServe(":3000", r)
}
