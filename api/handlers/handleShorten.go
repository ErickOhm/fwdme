package handlers

import (
	"fmt"
	"net/http"
)

type UrlStore interface {
	GetShortenedUrl(url string) string
	GetFullUrl(short string) string
}

type ShortenHandler struct {
	Store UrlStore
}

func (s *ShortenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.shortenUrl(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (s *ShortenHandler) shortenUrl(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	url := queryParams.Get("url")

	w.WriteHeader(http.StatusAccepted)

	short := s.Store.GetShortenedUrl(url)
	fmt.Fprint(w, short)

}
