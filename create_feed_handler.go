package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/friday1602/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func (a *apiConfig) createFeed(w http.ResponseWriter, r *http.Request, db database.User) {
	type reqParamsStruct struct {
		Name string `json:"name"`
		Url string `json:"url"`
	}
	reqParams := &reqParamsStruct{}
	err := json.NewDecoder(r.Body).Decode(reqParams)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "error decoding json")
		return 
	}

	feedParams := database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: reqParams.Name,
		Url: reqParams.Url,
		UserID: db.ID,
	}

	feed, err := a.DB.CreateFeed(context.Background(), feedParams)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "error decoding json")
		return 
	}

	respFeed := feedDatabaseToFeed(feed)
	responseWithJSON(w, http.StatusCreated, respFeed)
	

}
