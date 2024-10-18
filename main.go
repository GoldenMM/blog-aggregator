package main

import (
	"fmt"
	"log"

	"github.com/GoldenMM/blog-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("unable to read config: %v", err)
	}

	err = cfg.SetUser("ryan")
	if err != nil {
		log.Fatalf("unable to set user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("unable to read config: %v", err)
	}

	fmt.Println(cfg)
}
