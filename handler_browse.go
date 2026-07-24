package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/dhalvyr/blog_gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) >= 1 {
		val, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return err
		} else {
			limit = val
		}
	}

	posts, err := s.db.GetPostsForUser(
		context.Background(), 
		database.GetPostsForUserParams{
			UserID: user.ID,
			Limit: int32(limit),
		})
	if err != nil {
		return err
	}
	
	for _, post := range posts {
		fmt.Println(post)
	}

	return nil
}