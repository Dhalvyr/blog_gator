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

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("no name provided")
	}

	name := cmd.args[0]

	createdUser, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
	})
	if err != nil {
		fmt.Println("provided name already has an user")
		os.Exit(1)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return err
	}

	fmt.Printf("User %s successfully created.\n", name)
	log.Printf("%+v\n", createdUser)

	return nil
}