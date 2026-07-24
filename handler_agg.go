package main

import (
	"fmt"
	"os"
	"time"
)

func handlerAggregator(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		fmt.Println("time between requests required")
		os.Exit(1)
	}

	time_between_reqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %v\n", time_between_reqs)

	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}