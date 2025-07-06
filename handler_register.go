package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/greg-beach/blog_aggregator/sql/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]
	id := uuid.New()

	arg := database.CreateUserParams{
		id:         id,
		created_at: time.Now(),
		updated_at: time.Now(),
		name:       name,
	}

	user, err := dbQueries.GetUser(context.Background(), name)
	if err != nil {
		return err
	}
	if user {
		return fmt.Errorf("the user already exists")
	}
	
	_, err := dbQueries.CreateUser(context.Background(), arg)
	if err != nil {
		return err
	}

	return nil
}
