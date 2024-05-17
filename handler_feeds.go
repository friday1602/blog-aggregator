package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/friday1602/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

type reqParamsStruct struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (a *apiConfig) createFeed(w http.ResponseWriter, r *http.Request, db database.User) {

	reqParams := &reqParamsStruct{}
	err := json.NewDecoder(r.Body).Decode(reqParams)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "error encoding json")
		return
	}

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      reqParams.Name,
		Url:       reqParams.Url,
		UserID:    db.ID,
	}

	feed, err := a.DB.CreateFeed(r.Context(), feedParams)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "error creating feed")
		log.Println(err)
		return
	}

	feedFollows := database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		FeedID:    feedParams.ID,
		UserID:    db.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	createdFeedFollows, err := a.DB.CreateFeedFollows(r.Context(), feedFollows)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "error creating feed follows")
		return
	}

	respFeed := feedDatabaseToFeed(feed)
	respFeedFollows := feedFollowDatabaseToFeedFollow(createdFeedFollows)

	resp := FeedCreatedResp{
		Feed: respFeed,
		FeedFollows: respFeedFollows,
	}

	responseWithJSON(w, http.StatusCreated, resp)

}


func (a *apiConfig) getFeedsHandler(w http.ResponseWriter, r *http.Request) {
	feeds, err := a.DB.GetFeed(r.Context())
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "error getting feeds")
		return
	}
	responseWithJSON(w, http.StatusOK, feedsDatabaseToFeeds(feeds))

}