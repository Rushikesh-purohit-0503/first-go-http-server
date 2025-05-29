package main

import (
	"time"

	"github.com/Rushikesh-purohit-0503/rssagg/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	APIKey    string    `json:"api_key"`
}
type Feeds struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user._id"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		APIKey:    dbUser.ApiKey,
	}
}

func databaseFeedsToFeeds(dbfeeds database.Feed) Feeds {
	return Feeds{
		ID:        dbfeeds.ID,
		Name:      dbfeeds.Name,
		CreatedAt: dbfeeds.CreatedAt,
		UpdatedAt: dbfeeds.UpdatedAt,
		Url:       dbfeeds.Url,
		UserID:    dbfeeds.UserID,
	}
}

func databaseAllFeedsToAllFeeds(dbfeeds []database.Feed) []Feeds {
	feeds := []Feeds{}
	for _, dbfeed := range dbfeeds {
		feeds = append(feeds, databaseFeedsToFeeds(dbfeed))
	}
	return feeds
}
