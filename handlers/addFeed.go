package handlers

import (
	"context"
	"errors"
	"fmt"
	"gator/internal/database"
	"gator/shared"
	"time"

	"github.com/google/uuid"
)

func AddFeed(s *shared.State, cmd shared.Command, user database.User) error {
	if len(cmd.Args) != 2 {
		return errors.New("add-feed takes a Name and URL as arguments")
	}

	feed, err := s.Db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}

	_, err = s.Db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println(feed)
	return nil
}
