package main

import (
	"context"
	"fmt"
)

func handlerListFeeds(s *state, cmd command) error {
	// Check if the number of arguments is correct
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: feeds")
	}

	// Get the feeds from the database
	feed, err := s.db.ListFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("unable to list feeds: %v", err)
	}

	// Print the feeds
	for _, f := range feed {
		fmt.Printf("[%s]\n", f.Name)
		fmt.Printf("Created by: %s\n", f.UserName)
		fmt.Printf("  %s\n", f.Url)
	}

	return nil
}
