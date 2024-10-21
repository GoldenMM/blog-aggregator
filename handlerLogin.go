package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	// Check if the number of arguments is correct
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: login <username>")
	}

	// Check if the user exists
	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("user [%s] does not exist", cmd.args[0])
	}

	// Set the user
	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("unable to set user: %v", err)
	}
	fmt.Printf("User [%s] logged in\n", cmd.args[0])

	return nil
}
