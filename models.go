package main

import (
	"database/sql"
	"time"

	"github.com/friday1602/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

type FeedFollows struct {
	ID        uuid.UUID `json:"id"`
	FeedID    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FeedCreatedResp struct {
	Feed        Feed        `json:"feed"`
	FeedFollows FeedFollows `json:"feed_follow"`
}

type Post struct {
	ID          uuid.UUID      `json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Title       string         `json:"title"`
	Url         string         `json:"url"`
	Description sql.NullString `json:"description"`
	PublishedAt time.Time      `json:"published_at"`
	FeedID      uuid.UUID      `json:"feed_id"`
}

func userDatabaseToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	}
}

func feedDatabaseToFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    feed.UserID,
	}
}

func feedsDatabaseToFeeds(feeds []database.Feed) []Feed {
	newFeeds := make([]Feed, 0, len(feeds))
	for _, feed := range feeds {
		newFeeds = append(newFeeds, feedDatabaseToFeed(feed))
	}

	return newFeeds
}

func feedFollowDatabaseToFeedFollow(feedfollows database.FeedFollow) FeedFollows {
	return FeedFollows{
		ID:        feedfollows.ID,
		FeedID:    feedfollows.FeedID,
		UserID:    feedfollows.UserID,
		CreatedAt: feedfollows.CreatedAt,
		UpdatedAt: feedfollows.UpdatedAt,
	}

}

func feedFollowsDatabaseToFeedFollows(feedfollows []database.FeedFollow) []FeedFollows {
	newFeedfollows := make([]FeedFollows, 0, len(feedfollows))
	for _, feedfollow := range feedfollows {
		newFeedfollows = append(newFeedfollows, feedFollowDatabaseToFeedFollow(feedfollow))
	}

	return newFeedfollows
}

func postDatabaseToPost(post database.Post) Post {
	return Post{
		post.ID,
		post.CreatedAt,
		post.UpdatedAt,
		post.Title,
		post.Url,
		post.Description,
		post.PublishedAt,
		post.FeedID,
	}
}

func postsDatabaseToPosts(postsDB []database.Post) []Post {
	posts := make([]Post, 0, len(postsDB))
	for _, post := range postsDB {
		posts = append(posts, postDatabaseToPost(post))
	}
	return posts
}