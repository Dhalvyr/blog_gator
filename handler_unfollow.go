package main

import (
	"context"
	"fmt"
	"os"

	"github.com/dhalvyr/blog_gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		fmt.Println("required url of the feed to unfollow")
		os.Exit(1)
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = s.db.UnfollowFeed(context.Background(), database.UnfollowFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return nil
}