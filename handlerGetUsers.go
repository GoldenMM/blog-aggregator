package main

import (
	"context"
	"fmt"
)

func handlerGetUsers(s *state, cmd command) error {
	// Check if the number of arguments is correct
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: getUsers")
	}

	// Get the users
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("unable to get users: %v", err)
	}
	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s \n", user.Name)
		}
	}

	return nil
}
