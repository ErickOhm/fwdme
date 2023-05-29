package handlers

import (
	"net/http"
	"strings"
)

type RedirectHandler struct {
	Store UrlStore
}

func (re *RedirectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	short := strings.TrimPrefix(r.URL.Path, "/")
	full := re.Store.GetFullUrl(short)

	if full == "" {
		w.WriteHeader(http.StatusNotFound)
	} else {
		http.Redirect(w, r, re.Store.GetFullUrl(short), http.StatusFound)
	}

}
