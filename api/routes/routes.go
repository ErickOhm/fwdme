package routes

import (
	"fwdme/api/handlers"
	"fwdme/database/stores"
	"net/http"
)

func Routes() {
	http.Handle("/shorten", &handlers.ShortenHandler{Store: stores.NewInMemoryUrlStore()})
	http.Handle("/", &handlers.RedirectHandler{Store: stores.NewInMemoryUrlStore()})
}
