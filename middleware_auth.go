package main

import (
	"fmt"
	"net/http"

	"github.com/mekavehamichlolay/golang-course/internal/auth"
	"github.com/mekavehamichlolay/golang-course/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, "Coldn't get user")
			return
		}
		handler(w, r, user)
	}
}
