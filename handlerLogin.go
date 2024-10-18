package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	// Check if the number of arguments is correct
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: login <username>")
	}

	fmt.Println("Setting user:", cmd.args[0])

	// Set the user
	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("unable to set user: %v", err)
	}

	fmt.Println("User set to:", s.cfg.CurrentUserName)
	return nil
}
