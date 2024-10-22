package main

import (
	"context"
	"fmt"
	"time"

	"github.com/GoldenMM/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	// Check if the number of arguments is correct
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: follow <feed>")
	}

	// Check if feed exists
	var feed database.Feed
	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.args[0])
	if err != nil {
		// Not in database therefore fetch it from the internet
		rssFeed, err := fetchFeed(context.Background(), cmd.args[0])
		if err != nil {
			return fmt.Errorf("unable to fetch feed: %v", err)
		}
		// Add the feed
		now := time.Now()
		args := database.AddFeedParams{
			ID:        uuid.New(),
			Name:      rssFeed.Channel.Title,
			Url:       cmd.args[0],
			CreatedAt: now,
			UpdatedAt: now,
			UserID:    user.ID,
		}
		feed, err = s.db.AddFeed(context.Background(), args)
		if err != nil {
			return fmt.Errorf("unable to add feed: %v", err)
		}
	}
	// Create the link between the user and the feed
	now := time.Now()
	args := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		FeedID:    feed.ID,
		UserID:    user.ID,
	}
	row, err := s.db.CreateFeedFollow(context.Background(), args)
	if err != nil {
		return fmt.Errorf("unable to follow feed: %v", err)
	}
	fmt.Printf("User [%s] is now following feed [%s]\n", row.UserName, row.FeedName)
	return nil
}
