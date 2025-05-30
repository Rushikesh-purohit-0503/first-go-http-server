package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Rushikesh-purohit-0503/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:url`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	if err := decoder.Decode(&params); err != nil {
		responseWithError(w, 400, fmt.Sprintf("Invalid request payload : %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeeds(r.Context(), database.CreateFeedsParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Failed to create feed: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedsToFeeds(feed))
}

func (apiCfg *apiConfig) handlerGetAllFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Failed to get all feeds: %v", err))
		return
	}
	respondWithJSON(w, 201, databaseAllFeedsToAllFeeds(feeds))
}
