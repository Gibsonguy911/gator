package shared

import (
	"context"
	"errors"
	"gator/internal/config"
	"gator/internal/database"
)

type State struct {
	Config *config.Config
	Db     *database.Queries
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Commands map[string]func(*State, Command) error
}

func MiddlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, cmd Command) error {
		if s.Config.CurrentUserName == "" {
			return errors.New("no user logged in")
		}

		user, err := s.Db.GetUser(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return errors.New("no user logged in")
		}

		return handler(s, cmd, user)
	}
}

func (c *Commands) Register(name string, fn func(*State, Command) error) {
	c.Commands[name] = fn
}

func (c *Commands) Run(s *State, cmd Command) error {
	fn, ok := c.Commands[cmd.Name]
	if !ok {
		return errors.New("unknown command")
	}

	return fn(s, cmd)
}
