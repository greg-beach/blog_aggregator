package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/greg-beach/blog_aggregator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	dbQueries := s.db

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	arg := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	}

	_, err := dbQueries.GetUser(context.Background(), name)
	if err == nil {
		fmt.Fprintln(os.Stderr, "User already exists... exitting")
		os.Exit(1)
	}
	if err != sql.ErrNoRows {
		return err
	}

	user, err := dbQueries.CreateUser(context.Background(), arg)
	if err != nil {
		return err
	}

	s.cfg.SetUser(name)

	fmt.Println("user was created successfully:")
	printUser(user)

	return nil
}

func handlerLogin(s *state, cmd command) error {
	dbQueries := s.db

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	_, err := dbQueries.GetUser(context.Background(), name)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Username does not exist")
		os.Exit(1)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}
