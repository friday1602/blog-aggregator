package main

import "net/http"

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	responseWithJSON(w, http.StatusOK, struct{
		Status string `json:"status"`
	}{
		Status: "ok",
	})
}
