package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/friday1602/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func (a *apiConfig) createFeedFollows(w http.ResponseWriter, r *http.Request, db database.User) {
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
		ID:        uuid.New(),
		FeedID:    req.FeedID,
		UserID:    db.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	createdFeedFollows, err := a.DB.CreateFeedFollows(r.Context(), feedFollows)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "error creating feed follows")
		return
	}
	respFeedFollows := feedFollowDatabaseToFeedFollow(createdFeedFollows)
	responseWithJSON(w, http.StatusCreated, respFeedFollows)

}

func (a *apiConfig) deleteFeedFollow(w http.ResponseWriter, r *http.Request, userDB database.User) {
	feedFollowID := r.PathValue("feedFollowID")
	feedFollowUUID, err := uuid.Parse(feedFollowID)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, "error parsing feed follow id")
		return
	}
	err = a.DB.DeleteFeedFollow(r.Context(), feedFollowUUID)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "error deleting feed follow")
		return
	}

}

func (a *apiConfig) getFeedFollow(w http.ResponseWriter, r *http.Request, userDB database.User) {
	feedFollows, err := a.DB.GetFeedFollows(r.Context(), userDB.ID)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "error getting feed follows")
		return
	}
	feedFollowsResp := feedFollowsDatabaseToFeedFollows(feedFollows)
	responseWithJSON(w, http.StatusOK, feedFollowsResp)

}
