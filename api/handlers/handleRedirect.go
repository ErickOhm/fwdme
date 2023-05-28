package handlers

import (
	"net/http"
	"strings"
)

type RedirectStore interface {
	GetFullUrl(short string) string
}

type RedirectHandler struct {
	Store RedirectStore
}

func (re *RedirectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	short := strings.TrimPrefix(r.URL.Path, "/")
	http.Redirect(w, r, re.Store.GetFullUrl(short), http.StatusFound)
}
