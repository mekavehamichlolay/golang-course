package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	type status struct {
		Status string `json:"status"`
	}
	respondWithJSON(w, 200, status{Status: "ok"})
}
func handlerHealth(w http.ResponseWriter, r *http.Request) {

	respondWithJSON(w, 200, struct{}{})
}
