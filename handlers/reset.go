package handlers

import (
	"context"
	"fmt"
	"gator/shared"
)

func Reset(s *shared.State, cmd shared.Command) error {
	err := s.Db.ResetUsers(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("Deleted all users")

	err = s.Config.SetUser("")
	if err != nil {
		return err
	}

	fmt.Println("Removed current user from config")
	return nil
}
