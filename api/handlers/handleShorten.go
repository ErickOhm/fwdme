package handlers

import (
	"fmt"
	"net/http"
)

type ShortenStore interface {
	GetShortenedUrl(url string) string
}

type ShortenHandler struct {
	Store ShortenStore
}

func (s *ShortenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	url := queryParams.Get("url")

	fmt.Fprint(w, s.Store.GetShortenedUrl(url))
}
