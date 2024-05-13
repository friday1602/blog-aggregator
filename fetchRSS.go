package main

import (
	"context"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/friday1602/blog-aggregator/internal/database"
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
				rss, err := fetchRSS(feed.Name)
				if err != nil {
					log.Println("error fetching feeds:", err)
					return
				}

				// processing fetch data
				for _, item := range rss.Channel.Item {
					log.Println("Found Post", item.Title, "on feed", feed.Name)
				}

			}(feed)
		}
		wg.Wait()

	}

}
