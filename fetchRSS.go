package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/friday1602/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

type RSS struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Generator   string    `xml:"generator"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	Guid        string `xml:"guid"`
	Description string `xml:"description"`
}

func fetchRSS(url string) (RSS, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		return RSS{}, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return RSS{}, err
	}

	rss := RSS{}
	err = xml.Unmarshal(respBody, &rss)
	if err != nil {
		return RSS{}, err
	}

	return rss, nil

}

// scrap post and save to posts table on database
func worker(db *database.Queries, fetchInterval time.Duration, feedFetchedPerInterval int32) {
	// wait 60 sec interval
	interval := time.NewTicker(fetchInterval)
	for ; ; <-interval.C {
		// get the next feeds to fetch
		feeds, err := db.GetNextFeedsToFetch(context.Background(), feedFetchedPerInterval)
		if err != nil {
			log.Println("error getting feeds:", err)
			continue
		}

		// fetch from feed and mark feed fetched
		var wg sync.WaitGroup

		for _, feed := range feeds {
			wg.Add(1)

			go func(feed database.Feed) {
				defer wg.Done()
				// marking feed fetched
				err = db.MarkFeedFetch(context.Background(), feed.ID)
				if err != nil {
					log.Println("error marking feed as fetched:", err)
					return
				}

				// fetching each feed
				rss, err := fetchRSS(feed.Url)
				if err != nil {
					log.Println("error fetching feeds:", err)
					return
				}

				// processing fetch data
				for _, item := range rss.Channel.Item {
					pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)
					if err != nil {
						log.Println("error parsing pubdate:", err)
						continue
					}

					description := sql.NullString{}
					if item.Description != "" {
						description.String = item.Description
						description.Valid = true
					}
					_, err = db.CreatePost(context.Background(), database.CreatePostParams{
						ID: uuid.New(),
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
						Title: item.Title,
						Url: item.Link,
						Description: description,
						PublishedAt: pubDate,
						FeedID: feed.ID,
					})
					if err != nil {
						if strings.Contains(err.Error(), "duplicate key") {
							continue
						}
						log.Println("error creating post:", err)
					}
				}

			}(feed)
		}
		wg.Wait()

	}

}
