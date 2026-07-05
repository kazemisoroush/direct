// Package restaurant reads restaurants from storage and filters them by delivery area.
package restaurant

import (
	"context"

	"github.com/kazemisoroush/direct/backend/internal/domain"
)

//go:generate go tool mockgen -source=store.go -destination=../mocks/store_mock.go -package=mocks

// Store is the restaurant data access. ListDeliveringTo returns the delivery-area matches
// without menus; Get returns one restaurant with its full menu; Put creates or replaces one.
type Store interface {
	ListDeliveringTo(ctx context.Context, address string) ([]domain.Restaurant, error)
	Get(ctx context.Context, id string) (domain.Restaurant, error)
	Put(ctx context.Context, r domain.Restaurant) error
}
