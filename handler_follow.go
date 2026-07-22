package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/dhalvyr/blog_gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		fmt.Println("required a url to follow")
		os.Exit(1)
	}

	url := cmd.args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	createdFollow, err := s.db.CreateFeedFollow(
		context.Background(), 
		database.CreateFeedFollowParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID: user.ID,
			FeedID: feed.ID,

		})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(createdFollow.FeedName)
	fmt.Println(createdFollow.UserName)

	return nil
}