package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/friday1602/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func (a *apiConfig) createFeedFollows (w http.ResponseWriter, r *http.Request, db database.User) {
	type reqBody struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	
	req := reqBody{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, "error decoding json")
		return
	}

	feedFollows := database.CreateFeedFollowsParams{
		ID: uuid.New(),
		FeedID: req.FeedID,
		UserID: db.ID,
		CreatedAt: time.Now(),
	}
	createdFeedFollows, err := a.DB.CreateFeedFollows(r.Context(), feedFollows)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "error creating feed follows")
		return
	}

	respFeedFollows := feedFollowDatabaseToFeedFollow(createdFeedFollows)
	responseWithJSON(w, http.StatusCreated, respFeedFollows)

}