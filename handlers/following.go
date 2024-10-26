package handlers

import (
	"context"
	"fmt"
	"gator/internal/database"
	"gator/shared"
)

func Following(s *shared.State, _ shared.Command, user database.User) error {
	follows, err := s.Db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, follow := range follows {
		fmt.Println(follow.FeedName)
	}

	return nil
}
