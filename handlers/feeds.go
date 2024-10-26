package handlers

import (
	"context"
	"fmt"
	"gator/shared"
)

func Feeds(s *shared.State, cmd shared.Command) error {
	feeds, err := s.Db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("%s (%s) %s\n", feed.FeedName, feed.FeedUrl, feed.UserName)
	}

	return nil
}
