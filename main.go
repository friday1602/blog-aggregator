package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	mux := http.NewServeMux()

	mux.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Hello World!") })
	mux.HandleFunc("GET /v1/readiness", readinessHandler)
	mux.HandleFunc("GET /v1/err", errHandler)

	log.Println("starting server on port", port)
	err = http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}

}

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "failed to marshal json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)

}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	response, err := json.Marshal(struct{
		Error string `error:"json"`
	}{
		Error: msg,
	})
	if err != nil {
		http.Error(w, "failed to marshal json", http.StatusInternalServerError)
		return
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)


}
