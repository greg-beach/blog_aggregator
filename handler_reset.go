package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	dbQueries := s.db

	err := dbQueries.Reset(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("Database reset successfully!")

	return nil
}
