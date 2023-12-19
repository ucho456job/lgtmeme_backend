package main

import (
	handler "lgtmeme_backend/api"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.Handler)
	http.ListenAndServe(":3003", nil)
}
