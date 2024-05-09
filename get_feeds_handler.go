package main

import (
	"net/http"
)

func (a *apiConfig) getFeedsHandler(w http.ResponseWriter, r *http.Request) {
	feeds, err := a.DB.GetFeed(r.Context())
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "error getting feeds")
		return
	}
	responseWithJSON(w, http.StatusOK, feedsDatabaseToFeeds(feeds))

}
