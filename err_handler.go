package main

import "net/http"

func errHandler(w http.ResponseWriter, r *http.Request) {
	responseWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
