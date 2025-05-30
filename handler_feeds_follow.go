package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Rushikesh-purohit-0503/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	if err := decoder.Decode(&params); err != nil {
		responseWithError(w, 400, fmt.Sprintf("Invalid request payload : %v", err))
		return
	}

	FeedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Failed to create feed: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowstoFeedFollows(FeedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	FeedFollow, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Failed to Get Field: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseAllFeedsFollowToAllFeedsFollow(FeedFollow))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	followIDStr := chi.URLParam(r, "id")
	feedFollowID, err := uuid.Parse(followIDStr)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldn't parse uuid : %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Failed to unfollow: %v", err))
		return
	}
	respondWithJSON(w, 200, struct{}{})
}
