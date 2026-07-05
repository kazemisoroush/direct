package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/kazemisoroush/direct/backend/internal/domain"
	"github.com/kazemisoroush/direct/backend/internal/restaurant"
)

// RestaurantController serves the restaurant listing.
type RestaurantController struct {
	store restaurant.Store
}

// NewRestaurantController builds a RestaurantController over a store.
func NewRestaurantController(store restaurant.Store) *RestaurantController {
	return &RestaurantController{store: store}
}

// listRestaurantsResponse is the GET /restaurants body. Matches ListRestaurantsResponse in openapi.yaml.
type listRestaurantsResponse struct {
	Restaurants []domain.Restaurant `json:"restaurants"`
}

// List returns the restaurants that deliver to the address in the ?address= query.
func (c *RestaurantController) List(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	items, err := c.store.ListDeliveringTo(r.Context(), address)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "list restaurants")
		return
	}
	writeJSON(w, http.StatusOK, listRestaurantsResponse{Restaurants: items})
}

// Create stores a restaurant (with its menu), creating or replacing it. This is how
// restaurants enter the catalogue — API-first, no out-of-band writes to the store.
func (c *RestaurantController) Create(w http.ResponseWriter, r *http.Request) {
	var item domain.Restaurant
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "invalid restaurant body")
		return
	}
	if item.ID == "" || item.Name == "" {
		writeError(w, http.StatusBadRequest, "id and name are required")
		return
	}
	if err := c.store.Put(r.Context(), item); err != nil {
		writeError(w, http.StatusInternalServerError, "create restaurant")
		return
	}
	writeJSON(w, http.StatusCreated, item)
}

// Get returns one restaurant with its menu, or 404 when the id is unknown.
func (c *RestaurantController) Get(w http.ResponseWriter, r *http.Request) {
	restaurantItem, err := c.store.Get(r.Context(), r.PathValue("id"))
	if errors.Is(err, restaurant.ErrNotFound) {
		writeError(w, http.StatusNotFound, "restaurant not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "get restaurant")
		return
	}
	writeJSON(w, http.StatusOK, restaurantItem)
}
