package main

import (
	"net/http"

	"github.com/friday1602/blog-aggregator/internal/database"
)

func (a *apiConfig) getPostsByUser (w http.ResponseWriter, r *http.Request, userDB database.User) {
	postsDB, err := a.DB.GetPostByUser(r.Context(), database.GetPostByUserParams{
		UserID: userDB.ID,
		Limit: 10,
	})
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "error getting posts")
		return 
	}

	posts := postsDatabaseToPosts(postsDB)
	responseWithJSON(w, http.StatusOK, posts)
}