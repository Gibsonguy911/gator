package handlers

import (
	"context"
	"errors"
	"fmt"
	"gator/internal/database"
	"gator/shared"
)

func Unfollow(s *shared.State, cmd shared.Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return errors.New("unfollow takes a feed url as an argument")
	}

	err := s.Db.UnfollowFeed(context.Background(), database.UnfollowFeedParams{
		UserID: user.ID,
		Url:    cmd.Args[0],
	})
	if err != nil {
		return err
	}

	fmt.Println("Unfollowed", cmd.Args[0])

	return nil
}
