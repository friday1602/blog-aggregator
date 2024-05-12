package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
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
