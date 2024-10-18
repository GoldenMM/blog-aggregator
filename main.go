package main

import (
	"log"
	"os"

	"github.com/GoldenMM/blog-aggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	// Create the program's State
	c, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	s := &state{cfg: &c}

	// Create and register the commands
	cmds := commands{regCmds: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)

	// Read the commands from the input
	if len(os.Args) < 2 {
		log.Fatal("usage: blog-aggregator <command> [<args>]")
		return
	}
	cmd := command{name: os.Args[1], args: os.Args[2:]}
	err = cmds.run(s, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
