package main

import (
	"fmt"
	"log"

	"github.com/GoldenMM/blog-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	s := State{cfg: cfg}

	fmt.Println(s)
}
