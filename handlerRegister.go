package main

import (
	"context"
	"fmt"
	"time"

	"github.com/GoldenMM/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	// Check if the number of arguments is correct
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: register <username>")
	}

	// Check if the user already exists
	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err == nil {
		return fmt.Errorf("user [%s] already exists", cmd.args[0])
	}

	now := time.Now()
	args := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      cmd.args[0]}

	s.db.CreateUser(context.Background(), args)

	s.cfg.SetUser(cmd.args[0])
	fmt.Printf("User [%s] registered\n", cmd.args[0])

	return nil
}
