package main

import (
	"time"

	"github.com/friday1602/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
	ApiKey    string `json:"api_key"`
}


type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}


func userDatabaseToUser(user database.User) User {
	return User{
		ID: user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name: user.Name,
		ApiKey: user.ApiKey,
	}	
}

func feedDatabaseToFeed(feed database.Feed) Feed {
	return Feed{
		ID: feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name: feed.Name,
		Url: feed.Url,
		UserID: feed.UserID,
	}
}
