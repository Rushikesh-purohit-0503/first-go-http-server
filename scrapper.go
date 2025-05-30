package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/Rushikesh-purohit-0503/rssagg/internal/database"
	"github.com/google/uuid"
)

func startScrapping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Scrapping On %v goroutines every %s duration", concurrency, &timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Println("error fetching feeds:", err)
			continue
		}
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapFeed(db, wg, feed)
		}
		wg.Wait()
	}

}

func scrapFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("error making feed as fetched:", err)
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("Error fetching feed", err)
		return
	}
	for _, item := range rssFeed.Channel.Item {
		description := sql.NullString{}

		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		pubAt, err := time.Parse(time.RFC1123, item.PubDate)
		if err != nil {
			log.Printf("Error Parsing date %v with err %v", item.PubDate, err)
			continue
		}
		// log.Println("Post found:", item.Title , "On feed", feed.Name)
		 _ , err = db.CreatePosts(context.Background(), database.CreatePostsParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: description,
			PublishedAt: pubAt,
			FeedID:      feed.ID,
		})

		if err != nil {
			if strings.Contains(err.Error(),"duplicate key"){
				continue
			}
			log.Println("failed to create post:",err)
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
