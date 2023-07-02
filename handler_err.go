package main

import "net/http"

func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 200, "something went wrong")
}
