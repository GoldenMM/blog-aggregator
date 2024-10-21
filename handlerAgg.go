package main

import (
	"context"
	"fmt"
)

func handlerAgg(_ *state, cmd command) error {
	// Check if the number of arguments is correct
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: agg")
	}

	// Run the rss feed
	const testUrl = "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), testUrl)
	if err != nil {
		return fmt.Errorf("unable to fetch feed: %v", err)

	}

	// Print the feed
	fmt.Println(feed)
	return nil

}
