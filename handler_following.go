package main

import (
	"context"
	"fmt"
	"os"

	"github.com/dhalvyr/blog_gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
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