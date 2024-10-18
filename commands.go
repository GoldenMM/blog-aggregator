package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	regCmds map[string]func(*state, command) error
}

func (c *commands) register(name string, handler func(*state, command) error) {
	c.regCmds[name] = handler
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.regCmds[cmd.name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.name)
	}
	return handler(s, cmd)
}
