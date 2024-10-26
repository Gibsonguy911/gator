package handlers

import (
	"context"
	"errors"
	"fmt"
	"gator/shared"
)

func Login(s *shared.State, cmd shared.Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("login takes a name as an argument")
	}

	_, err := s.Db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("%s not registered", cmd.Args[0])
	}

	err = s.Config.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Println("Logged in as", cmd.Args[0])
	return nil
}
