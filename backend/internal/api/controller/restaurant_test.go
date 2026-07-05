package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kazemisoroush/direct/backend/internal/domain"
)

// fakeStore returns a fixed set of restaurants (or an error).
type fakeStore struct {
	items []domain.Restaurant
	err   error
}

func (f fakeStore) ListDeliveringTo(_ context.Context, _ string) ([]domain.Restaurant, error) {
	return f.items, f.err
}

func TestListReturnsEmptyArray(t *testing.T) {
	// Arrange: no restaurants seeded yet (M1).
	c := NewRestaurantController(fakeStore{items: []domain.Restaurant{}})
	req := httptest.NewRequest(http.MethodGet, "/restaurants?address=Kellyville+NSW+2155", nil)
	rec := httptest.NewRecorder()

	// Act
	c.List(rec, req)

	// Assert: 200 with an explicit empty array, not null.
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	var body struct {
		Restaurants []domain.Restaurant `json:"restaurants"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode body: %v", err)
	}
	if body.Restaurants == nil {
		t.Fatal("restaurants is null, want []")
	}
	if len(body.Restaurants) != 0 {
		t.Fatalf("len = %d, want 0", len(body.Restaurants))
	}
}

func TestListReturnsRestaurants(t *testing.T) {
	c := NewRestaurantController(fakeStore{items: []domain.Restaurant{{ID: "1", Name: "Hills Kebabs"}}})
	req := httptest.NewRequest(http.MethodGet, "/restaurants", nil)
	rec := httptest.NewRecorder()

	c.List(rec, req)

	var body struct {
		Restaurants []domain.Restaurant `json:"restaurants"`
	}
	_ = json.Unmarshal(rec.Body.Bytes(), &body)
	if len(body.Restaurants) != 1 || body.Restaurants[0].Name != "Hills Kebabs" {
		t.Fatalf("body = %+v, want one Hills Kebabs", body.Restaurants)
	}
}

func TestListSurfacesStoreError(t *testing.T) {
	c := NewRestaurantController(fakeStore{err: context.DeadlineExceeded})
	req := httptest.NewRequest(http.MethodGet, "/restaurants", nil)
	rec := httptest.NewRecorder()

	c.List(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d, want 500 on store error", rec.Code)
	}
}
