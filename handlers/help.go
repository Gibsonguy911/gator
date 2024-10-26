package handlers

import (
	"fmt"
	"gator/shared"
)

func printMessage(cmd, desc string, args []string) {
	argPrint := ""
	for _, arg := range args {
		argPrint += fmt.Sprintf("<%s> ", arg)
	}
	fmt.Println(" ", cmd)
	fmt.Println("    ", desc)
	if len(args) > 0 {
		fmt.Println("      Args:", argPrint)
	}
}

func Help(s *shared.State, cmd shared.Command) error {
	fmt.Println("Available commands:")

	printMessage("register", "Register a new user. Updates config to set current user.", []string{"username"})
	printMessage("login", "Log in to the application. User must be registered. Updates config to set current user.", []string{"username"})
	printMessage("reset", "Delete all users and remove current user from config.", []string{})
	printMessage("addFeed", "Add a new RSS Feed to the application. User must be logged in.", []string{"feedUrl"})
	printMessage("feeds", "List all feeds in the application.", []string{})
	printMessage("follow", "Follow a feed. User must be logged in.", []string{"feedUrl"})
	printMessage("following", "List all feeds the user is following. User must be logged in.", []string{})
	printMessage("unfollow", "Unfollow a feed. User must be logged in.", []string{"feedUrl"})
	printMessage("browse", "Browse all articles from all feeds the user is following. User must be logged in.", []string{"limit (default: 2)"})
	printMessage("users", "List all users in the application.", []string{})
	printMessage("agg", "Aggregate all articles from all feeds the user is following. Constantly runs in the terminal session on a timer. User must be logged in.", []string{"refresh timer e.g. 1m, 1h, 1d"})
	return nil
}
