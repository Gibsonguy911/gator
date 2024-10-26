package handlers

import (
	"context"
	"fmt"
	"gator/shared"
)

func Users(s *shared.State, cmd shared.Command) error {
	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		line := fmt.Sprintf("* %s", user.Name)
		if s.Config.CurrentUserName == user.Name {
			line += " (current)"
		}
		fmt.Println(line)
	}

	return nil
}
