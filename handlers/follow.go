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

func Follow(s *shared.State, cmd shared.Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return errors.New("follow takes a feed url as an argument")
	}

	feed, err := s.Db.GetFeedByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}

	feedFollow, err := s.Db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	fmt.Println("Followed", feedFollow.FeedName)

	return nil
}
