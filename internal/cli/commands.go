package cli

import (
	"fmt"
)

type commands struct {
	cmds map[string]func(*State, command) error
}

func (c *commands) register(name string, handler func(*State, command) error) {
	c.cmds[name] = handler
}

func (c *commands) run(s *State, cmd command) error {
	handler, ok := c.cmds[cmd.name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.name)
	}
	return handler(s, cmd)
}

type command struct {
	name string
	args []string
}

func handlerLogin(s *State, cmd command) error {
	// Check if the number of arguments is correct
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: login <username>")
	}

	// Set the user
	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("unable to set user: %v", err)
	}
	return nil
}
