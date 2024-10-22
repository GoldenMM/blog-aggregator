package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	// Check if the number of arguments is correct
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: following")
	}

	// Get the feeds from the database
	feed_following, err := s.db.GetFeedFollowsForUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("unable to list feeds: %v", err)
	}

	// Print the feeds
	for _, f := range feed_following {
		fmt.Printf("%s\n", f.FeedName)
	}

	return nil
}
