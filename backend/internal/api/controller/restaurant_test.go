package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/kazemisoroush/direct/backend/internal/domain"
	"github.com/kazemisoroush/direct/backend/internal/mocks"
)

func TestListReturnsEmptyArray(t *testing.T) {
	// Arrange: no restaurants seeded yet (M1).
	ctrl := gomock.NewController(t)
	store := mocks.NewMockStore(ctrl)
	store.EXPECT().ListDeliveringTo(gomock.Any(), gomock.Any()).Return([]domain.Restaurant{}, nil)
	req := httptest.NewRequest(http.MethodGet, "/restaurants?address=Kellyville+NSW+2155", nil)
	rec := httptest.NewRecorder()

	// Act
	NewRestaurantController(store).List(rec, req)

	// Assert: 200 with an explicit empty array on the wire, not null.
	require.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), `"restaurants":[]`)
}

func TestListReturnsRestaurants(t *testing.T) {
	ctrl := gomock.NewController(t)
	store := mocks.NewMockStore(ctrl)
	store.EXPECT().ListDeliveringTo(gomock.Any(), gomock.Any()).
		Return([]domain.Restaurant{{ID: "1", Name: "Hills Kebabs", Suburb: "Kellyville"}}, nil)
	req := httptest.NewRequest(http.MethodGet, "/restaurants", nil)
	rec := httptest.NewRecorder()

	NewRestaurantController(store).List(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)
	var body struct {
		Restaurants []domain.Restaurant `json:"restaurants"`
	}
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &body))
	require.Len(t, body.Restaurants, 1)
	assert.Equal(t, "Hills Kebabs", body.Restaurants[0].Name)
}

func TestListSurfacesStoreError(t *testing.T) {
	ctrl := gomock.NewController(t)
	store := mocks.NewMockStore(ctrl)
	store.EXPECT().ListDeliveringTo(gomock.Any(), gomock.Any()).
		Return(nil, errors.New("dynamo down"))
	req := httptest.NewRequest(http.MethodGet, "/restaurants", nil)
	rec := httptest.NewRecorder()

	NewRestaurantController(store).List(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}
