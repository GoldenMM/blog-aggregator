package main

import (
	"context"
	"fmt"

	"github.com/GoldenMM/blog-aggregator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	// Check if the number of arguments is correct
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: unfollow <feed>")
	}

	// Get the feed
	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("unable to get feed [%s]: %v", cmd.args[0], err)
	}

	args := database.DeleteFeedFollowParams{
		Name: user.Name,
		Url:  feed.Url,
	}
	// Unfollow the feed
	err = s.db.DeleteFeedFollow(context.Background(), args)
	if err != nil {
		return fmt.Errorf("unable to unfollow feed: %v", err)
	}
	fmt.Printf("User [%s] is no longer following feed [%s]\n", user.Name, feed.Name)
	return nil
}
