package handlers

import (
	"context"
	"errors"
	"fmt"
	"gator/internal/database"
	"gator/internal/rss"
	"gator/shared"
	"time"

	"github.com/araddon/dateparse"
	"github.com/google/uuid"
)

func scrapeFeeds(s *shared.State) error {
	feed, err := s.Db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("Fetching", feed.Url)
	feedData, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	for _, item := range feedData.Channel.Item {
		date, err := dateparse.ParseAny(item.PubDate)
		if err != nil {
			return err
		}

		_, err = s.Db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			PublishedAt: date,
			FeedID:      feed.ID,
		})

		if err != nil {
			fmt.Println("Error creating post", err)
			continue
		}
	}

	return nil
}

func Agg(s *shared.State, cmd shared.Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("agg takes a time as an argument. e.g. 1s, 1m, 1h")
	}

	d, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}

	ticker := time.NewTicker(d)
	for ; ; <-ticker.C {
		fmt.Printf("Fetching feeds every %s\n", d)
		scrapeFeeds(s)
	}
}
