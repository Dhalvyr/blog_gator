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

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("feed name and url required")
	}

	name := cmd.args[0]
	url := cmd.args[1]
	username := s.cfg.CurrentUserName
	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	user_id := user.ID

	createdFeed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
		Url: url,
		UserID: user_id,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	log.Printf("%+v\n", createdFeed)
	return nil
}