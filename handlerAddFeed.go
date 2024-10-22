package main

import (
	"context"
	"fmt"
	"time"

	"github.com/GoldenMM/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	// Check if the number of arguments is correct
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: addFeed <name> <url>")
	}

	// Add the feed
	now := time.Now()
	args := database.AddFeedParams{
		ID:        uuid.New(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
	}

	feed, err := s.db.AddFeed(context.Background(), args)
	if err != nil {
		return fmt.Errorf("unable to add feed: %v", err)
	}

	// Add the user to be following the feed
	args2 := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		FeedID:    feed.ID,
		UserID:    user.ID,
		CreatedAt: now,
		UpdatedAt: now,
	}
	_, err = s.db.CreateFeedFollow(context.Background(), args2)
	if err != nil {
		return fmt.Errorf("unable to follow feed: %v", err)
	}
	fmt.Printf("Feed [%s] added\n", feed.Name)

	return nil
}
