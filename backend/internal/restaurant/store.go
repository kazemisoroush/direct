// Package restaurant reads restaurants from storage and filters them by delivery area.
package restaurant

import (
	"context"

	"github.com/kazemisoroush/direct/backend/internal/domain"
)

//go:generate go tool mockgen -source=store.go -destination=../mocks/store_mock.go -package=mocks

// Store lists restaurants that deliver to a given address.
type Store interface {
	ListDeliveringTo(ctx context.Context, address string) ([]domain.Restaurant, error)
}
