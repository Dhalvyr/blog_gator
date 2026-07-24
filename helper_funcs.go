package main

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"log"

	"github.com/dhalvyr/blog_gator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return err
	}

	RSSFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	for _, post := range RSSFeed.Channel.Item {
		publishedAt := sql.NullTime{
			Valid: false,
		}
		parsedTime, err := time.Parse(time.RFC1123, post.PubDate)
		if err == nil {
			publishedAt.Valid = true
			publishedAt.Time = parsedTime
		}

		title := sql.NullString{
			String: post.Title,
			Valid: post.Title != "",
		}

		description := sql.NullString{
			String: post.Description,
			Valid: post.Description != "",
		}

		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: title,
			Url: post.Link,
			Description: description,
			PublishedAt: publishedAt,
			FeedID: feed.ID,
		})
		if err != nil {
			var pqErr *pq.Error
			if errors.As(err, &pqErr) && pqErr.Code == "23505" {
				continue
			} else {
					log.Println(err)
				}
		}
	}

	return nil
}