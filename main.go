package main

import (
	handler "LGTMeme_backend/api"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.Handler)
	http.ListenAndServe(":3003", nil)
}
