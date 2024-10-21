package main

import (
	"context"
	"fmt"
	"time"

	"github.com/GoldenMM/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	// Check if the number of arguments is correct
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: addFeed <name> <url>")
	}

	// Get current user
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("unable to get user: %v", err)
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
	_, err = s.db.AddFeed(context.Background(), args)
	if err != nil {
		return fmt.Errorf("unable to add feed: %v", err)
	}

	return nil
}
