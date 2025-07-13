package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/greg-beach/blog_aggregator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]

	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("unable to get feed: %w", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("feed could not be followed: %w", err)
	}

	fmt.Println("Feed has been followed successfully:")
	fmt.Printf("feed:     %s\n", feed.Name)
	fmt.Printf("user:     %s\n", user.Name)

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("unable to get feeds user is following: %w", err)
	}

	fmt.Printf("user: %s is following: \n", user.Name)

	for _, feed := range feeds {
		fmt.Printf("%s\n", feed.FeedName)
	}

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]

	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("unable to get feed: %w", err)
	}

	err = s.db.RemoveFeedFollow(context.Background(), database.RemoveFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't remove feed follow: %w", err)
	}

	fmt.Printf("%s unfollowed successfully!\n", feed.Name)
	return nil
}
