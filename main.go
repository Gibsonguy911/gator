package main

import (
	"database/sql"
	"fmt"
	"gator/handlers"
	"gator/internal/config"
	"gator/internal/database"
	"gator/shared"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// read config
	conf, err := config.Read()
	if err != nil {
		panic(err)
	}

	// connect to database
	db, err := sql.Open("postgres", conf.DbUrl)
	if err != nil {
		fmt.Println("Error connecting to datbase")
		os.Exit(1)
	}

	// create state
	dbQueries := database.New(db)
	s := shared.State{
		Config: &conf,
		Db:     dbQueries,
	}

	// register commands
	cmds := shared.Commands{Commands: make(map[string]func(*shared.State, shared.Command) error)}
	cmds.Register("help", handlers.Help)
	cmds.Register("login", handlers.Login)
	cmds.Register("register", handlers.Resiger)
	cmds.Register("reset", handlers.Reset)
	cmds.Register("users", handlers.Users)
	cmds.Register("agg", handlers.Agg)
	cmds.Register("addfeed", shared.MiddlewareLoggedIn(handlers.AddFeed))
	cmds.Register("feeds", handlers.Feeds)
	cmds.Register("follow", shared.MiddlewareLoggedIn(handlers.Follow))
	cmds.Register("following", shared.MiddlewareLoggedIn(handlers.Following))
	cmds.Register("unfollow", shared.MiddlewareLoggedIn(handlers.Unfollow))
	cmds.Register("browse", shared.MiddlewareLoggedIn(handlers.Browse))

	// parse user command
	if len(os.Args) < 2 {
		fmt.Println("Usage: gator <command> [args]")
		os.Exit(1)
	}

	commandName := os.Args[1]
	args := os.Args[2:]
	cmd := shared.Command{Name: commandName, Args: args}

	// run command
	err = cmds.Run(&s, cmd)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
