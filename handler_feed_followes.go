package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/mekavehamichlolay/golang-course/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json: %v", err))
		return
	}
	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		fmt.Println("Coldn't create feed follow because:", err)
		respondWithError(w, 400, "Coldn't create feed follow")
		return
	}
	respondWithJSON(w, 201, databaseFeedFollowesToFeedFollowes(feedFollow))

}

func (apiCfg *apiConfig) handlerGetFeedFollowes(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollowes, err := apiCfg.DB.GetFeedFollowes(r.Context(), user.ID)
	if err != nil {
		fmt.Println("Coldn't get feed followes:", err)
		respondWithError(w, 400, "Coldn't get feeds")
		return
	}
	respondWithJSON(w, 200, databaseFeedsFollowToFeedsFollow(feedFollowes))

}
func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		fmt.Printf("couldn't parse feed id: %v", err)
		respondWithError(w, 400, "coldn't parse feed id")
		return
	}
	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		fmt.Printf("couldn't delete feed follow: %v", err)
		respondWithError(w, 400, "coldn't delete feed follow")
		return
	}
	respondWithJSON(w, 201, struct{}{})
}
