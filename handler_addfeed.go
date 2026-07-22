package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dhalvyr/blog_gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("feed name and url required")
	}

	name := cmd.args[0]
	url := cmd.args[1]

	createdFeed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
		Url: url,
		UserID: user.ID,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = s.db.CreateFeedFollow(
		context.Background(), 
		database.CreateFeedFollowParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID: user.ID,
			FeedID: createdFeed.ID,

		})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	log.Printf("%+v\n", createdFeed)
	return nil
}