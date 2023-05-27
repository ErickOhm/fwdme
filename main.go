package main

import (
	"fwdme/api/routes"
	"log"
	"net/http"
)

func main() {
	routes.RegisterRoutes()
	log.Fatal(http.ListenAndServe(":5000", nil))
}
