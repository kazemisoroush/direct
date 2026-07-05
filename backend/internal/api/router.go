// Package api assembles the HTTP router, controllers, and middleware.
package api

import (
	"net/http"

	"github.com/kazemisoroush/direct/backend/internal/api/controller"
)

// Router maps HTTP endpoints to their controllers.
type Router struct {
	mux *http.ServeMux
}

// NewRouter wires each endpoint to a controller method.
func NewRouter(restaurants *controller.RestaurantController, health *controller.HealthController) *Router {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /restaurants", restaurants.List)
	mux.Handle("GET /health", health)
	return &Router{mux: mux}
}

// ServeHTTP dispatches a request to the matching controller.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
