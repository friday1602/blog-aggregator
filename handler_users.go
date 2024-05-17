package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/friday1602/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

type user struct {
	Name string `json:"name"`
}

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (a *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user user
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Error decoding json")
		return
	}

	db, err := a.DB.CreateUser(r.Context(), database.CreateUserParams{
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

func getUserHandler(w http.ResponseWriter, r *http.Request, db database.User) {
	userDB := userDatabaseToUser(db)
	responseWithJSON(w, http.StatusOK, userDB)

}
