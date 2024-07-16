package main

import (
	"encoding/json"
	"net/http"
)

func responseWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "failed to marshal json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)

}

func responseWithError(w http.ResponseWriter, code int, msg string) {
	response, err := json.Marshal(struct {
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
