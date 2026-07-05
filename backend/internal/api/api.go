// Package api wires the Direct HTTP routes. M0 exposes only a health probe; product
// routes (auth, restaurants, menu, orders) land in later milestones.
package api

import (
	"encoding/json"
	"net/http"
)

// New returns the HTTP handler serving the Direct API.
func New() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", health)
	return mux
}

// health is an unauthenticated liveness probe.
func health(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
