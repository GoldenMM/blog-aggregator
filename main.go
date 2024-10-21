// "postgres://postgres:postgres@localhost:5432/gator"

package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/GoldenMM/blog-aggregator/internal/config"
	"github.com/GoldenMM/blog-aggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {

	// Create the program's State
	c, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	// Open a connection to the database
	db, err := sql.Open("postgres", c.DbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	s := &state{db: dbQueries, cfg: &c}

	// Create and register the commands
	cmds := commands{regCmds: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)

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
