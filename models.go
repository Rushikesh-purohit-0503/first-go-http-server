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
	UserID    uuid.UUID `json:"user_id"`
}

type FeedFollows struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

type Posts struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
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

func databaseFeedFollowstoFeedFollows(dbFeedFollow database.FeedFollow) FeedFollows {
	return FeedFollows{
		ID:        dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		UserID:    dbFeedFollow.UserID,
		FeedID:    dbFeedFollow.FeedID,
	}
}

func databaseAllFeedsFollowToAllFeedsFollow(dbfeeds []database.FeedFollow) []FeedFollows {
	FeedFollow := []FeedFollows{}
	for _, dbfeed := range dbfeeds {
		FeedFollow = append(FeedFollow, databaseFeedFollowstoFeedFollows(dbfeed))
	}
	return FeedFollow
}

func databasePostToPost(dbPosts database.Post) Posts {
	var description *string
	if dbPosts.Description.Valid {
		description = &dbPosts.Description.String
	}
	return Posts{
		ID:          dbPosts.ID,
		CreatedAt:   dbPosts.CreatedAt,
		UpdatedAt:   dbPosts.UpdatedAt,
		FeedID:      dbPosts.FeedID,
		Title:       dbPosts.Title,
		Url:         dbPosts.Url,
		Description: description,
		PublishedAt: dbPosts.PublishedAt,
	}
}

func databasePostsToPosts(dbPosts []database.Post) []Posts {
	posts := []Posts{}

	for _, dbPost := range dbPosts {
		posts = append(posts, databasePostToPost(dbPost))
	}

	return posts
}
