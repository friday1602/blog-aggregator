package main

import (
	"net/http"

	"github.com/friday1602/blog-aggregator/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func getUserHandler(w http.ResponseWriter, r *http.Request, db database.User) {
	userDB := userDatabaseToUser(db)
	responseWithJSON(w, http.StatusOK, userDB)

}
