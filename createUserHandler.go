package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/friday1602/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

type user struct {
	Name string `json:"name"`
}

func (a *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user user
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Error decoding json")
		return
	}

	db, err := a.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      user.Name,
	})
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Error creating user")
		return 
	}

	responseWithJSON(w, http.StatusCreated, db)




}
