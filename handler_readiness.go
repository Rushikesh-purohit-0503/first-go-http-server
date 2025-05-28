package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}

func handleErrors(w http.ResponseWriter, r *http.Request) {
	responseWithError(w, 500, "Internal Server Error")
}