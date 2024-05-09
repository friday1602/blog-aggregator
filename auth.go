package main

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

func getAPIkey(header http.Header) (string, error) {
	apiFromHeader := header.Get("Authorization")
	apiParts := strings.Split(apiFromHeader, " ")

	if len(apiParts) != 2 || apiParts[0] != "ApiKey" {
		return "", errors.New("invalid apikey")
	}
	return apiParts[1], nil

}

func (a *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := getAPIkey(r.Header)
		if err != nil {
			responseWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		db, err := a.DB.GetUserByAPI(context.Background(), apiKey)
		if err != nil {
			responseWithError(w, http.StatusUnauthorized, "unauthorized")
			return
		}

		handler(w, r, db)
	}
}
