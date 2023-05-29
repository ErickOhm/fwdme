package main

import (
	"fwdme/api/routes"
	"log"
	"net/http"
)

func main() {
	routes.Routes()
	log.Fatal(http.ListenAndServe(":5300", nil))
}
