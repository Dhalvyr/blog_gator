package main

import (
	"context"
	"fmt"
	"os"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("no arguments required for this command")
	}

	data, err := s.db.GetFeeds(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := range data {
		feedName := data[i].FeedName
		url := data[i].Url
		userName := data[i].UserName
		fmt.Printf("Feed Name: %s url: %s created by user: %s.\n", feedName, url, userName)
	}
	return nil
}