// Package restaurant reads restaurants from storage and filters them by delivery area.
package restaurant

import (
	"context"

	"github.com/kazemisoroush/direct/backend/internal/domain"
)

//go:generate go tool mockgen -source=store.go -destination=../mocks/store_mock.go -package=mocks

// Store reads restaurants. ListDeliveringTo returns the delivery-area matches without menus;
// Get returns one restaurant with its full menu.
type Store interface {
	ListDeliveringTo(ctx context.Context, address string) ([]domain.Restaurant, error)
	Get(ctx context.Context, id string) (domain.Restaurant, error)
}
