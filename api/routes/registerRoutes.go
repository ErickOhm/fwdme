package routes

import "net/http"

func RegisterRoutes() {
	http.HandleFunc("/shorten", Shorten)
	http.HandleFunc("/", Redirect)
}
