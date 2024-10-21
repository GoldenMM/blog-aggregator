package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	// Check if the number of arguments is correct
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: reset")
	}

	// Reset the user
	err := s.db.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("unable to reset users table: %v", err)
	}
	fmt.Println("Users table reset")

	return nil
}
