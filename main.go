package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/friday1602/blog-aggregator/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DSN")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	newDB := database.New(db)
	apiCfg := apiConfig{
		DB: newDB,
	}

	go worker(newDB, 60*time.Second, 10)

	port := os.Getenv("PORT")
	mux := http.NewServeMux()

	mux.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Hello World!") })
	mux.HandleFunc("GET /v1/readiness", readinessHandler)
	mux.HandleFunc("GET /v1/err", errHandler)
	mux.HandleFunc("POST /v1/users", apiCfg.createUserHandler)
	mux.HandleFunc("GET /v1/users", apiCfg.middlewareAuth(getUserHandler))
	mux.HandleFunc("POST /v1/feeds", apiCfg.middlewareAuth(apiCfg.createFeed))
	mux.HandleFunc("GET /v1/feeds", apiCfg.getFeedsHandler)
	mux.HandleFunc("POST /v1/feed_follows", apiCfg.middlewareAuth(apiCfg.createFeedFollows))
	mux.HandleFunc("DELETE /v1/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.deleteFeedFollow))
	mux.HandleFunc("GET /v1/feed_follows", apiCfg.middlewareAuth(apiCfg.getFeedFollow))
	mux.HandleFunc("GET /v1/posts", apiCfg.middlewareAuth(apiCfg.getPostsByUser))
	srv := http.Server{
		Addr:              port,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("starting server on port", port)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
