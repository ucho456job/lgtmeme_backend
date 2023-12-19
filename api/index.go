package api

import (
	"net/http"
)

// This function is the entry point for vercel's server less function.
func Handler(w http.ResponseWriter, r *http.Request) {
	engine.ServeHTTP(w, r)
}
