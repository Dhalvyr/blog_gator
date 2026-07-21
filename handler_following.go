package main

import (
	"context"
	"fmt"
	"os"
)

func handlerFollowing(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, follow := range follows {
		fmt.Println(follow.FeedName)
	}

	return nil
}