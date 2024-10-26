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

func Resiger(s *shared.State, cmd shared.Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("register takes a name as an argument")
	}

	name := cmd.Args[0]
	user, err := s.Db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	err = s.Config.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Println("ID:", user.ID)
	fmt.Println("Name:", user.Name)
	fmt.Println("Created At:", user.CreatedAt)
	fmt.Println("Updated At:", user.UpdatedAt)
	return nil
}
