// Package restaurant reads restaurants from storage and filters them by delivery area.
package restaurant

import (
	"context"

	"github.com/kazemisoroush/direct/backend/internal/domain"
)

// Store lists restaurants that deliver to a given address.
type Store interface {
	ListDeliveringTo(ctx context.Context, address string) ([]domain.Restaurant, error)
}
