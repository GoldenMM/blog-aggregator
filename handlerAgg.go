package main

import (
	"context"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	// Check if the number of arguments is correct
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: agg <time_between_requests>")
	}
	// Parse the time between requests
	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("unable to parse time between requests: %v", err)
	}
	fmt.Printf("Aggregating feeds every %s\n", timeBetweenRequests)
	// Create a ticker
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			fmt.Printf("error scraping feeds: %v\n", err)
		}
	}
}

func scrapeFeeds(s *state) error {
	// Get the next feed to scrape
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("unable to get next feed to fetch: %v", err)
	}
	// Mark as fetched
	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("unable to mark feed as fetched: %v", err)
	}
	// Fetch the feed
	rrsFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("unable to fetch feed: %v", err)
	}
	for _, item := range rrsFeed.Channel.Items {
		fmt.Printf("Title: %s\n", item.Title)
	}
	return nil
}
