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

	respFeed := feedDatabaseToFeed(feed)
	responseWithJSON(w, http.StatusCreated, respFeed)

}
