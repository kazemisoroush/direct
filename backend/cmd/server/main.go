// Local HTTP server for running the Direct API outside Lambda during development.
package main

import (
	"log"
	"net/http"

	"github.com/kazemisoroush/direct/backend/internal/api"
)

func main() {
	const addr = ":8080"
	log.Printf("Direct API listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, api.New()))
}
