package handlers

import (
	"context"
	"fmt"
	"gator/internal/database"
	"gator/shared"
	"strconv"
)

func Browse(s *shared.State, cmd shared.Command, user database.User) error {
	limit := int32(2)
	if len(cmd.Args) == 1 {
		l, err := strconv.ParseInt(cmd.Args[0], 10, 32)
		if err != nil {
			return err
		}
		limit = int32(l)
	}

	posts, err := s.Db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	})
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Println(post.Title)
	}

	return nil
}
