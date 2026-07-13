package main

import (
	"context"
	"fmt"
	"os"
)

func handlerUsers(s *state, cmd command) error {
	data, err := s.db.GetUsers(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, user := range data {
		name := user.Name
		if name == s.cfg.CurrentUserName {
			name = name + " (current)"
		}
		fmt.Printf("* %s\n",name)
	}

	return nil
}