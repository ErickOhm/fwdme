package routes

import (
	"fwdme/api/handlers"
	"net/http"
)

func Shorten(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handler := &handlers.ShortenHandler{}
		handler.ServeHTTP(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

}
