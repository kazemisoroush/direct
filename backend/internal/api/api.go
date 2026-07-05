package api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/kazemisoroush/direct/backend/internal/api/controller"
	"github.com/kazemisoroush/direct/backend/internal/api/middleware"
	"github.com/kazemisoroush/direct/backend/internal/auth"
	"github.com/kazemisoroush/direct/backend/internal/config"
	"github.com/kazemisoroush/direct/backend/internal/restaurant"
)

// New builds the API handler: controllers behind the router, wrapped in middleware.
func New(ctx context.Context, cfg config.Config, store restaurant.Store) (http.Handler, error) {
	router := NewRouter(
		controller.NewRestaurantController(store),
		controller.NewHealthController(),
	)

	authed, err := authenticate(ctx, cfg, router)
	if err != nil {
		return nil, fmt.Errorf("configure authentication: %w", err)
	}
	return middleware.NewRecoverMiddleware().Wrap(authed), nil
}

// authenticate wraps the router with JWT auth, failing closed unless opted out.
func authenticate(ctx context.Context, cfg config.Config, routes http.Handler) (http.Handler, error) {
	if cfg.AuthDisabled {
		log.Print("auth explicitly disabled via DIRECT_AUTH_DISABLED; serving without authentication")
		return routes, nil
	}
	if !cfg.AuthEnabled() {
		return nil, errors.New("auth not configured: set DIRECT_JWT_ISSUER and DIRECT_JWT_CLIENT_ID, or set DIRECT_AUTH_DISABLED=true to run without auth")
	}

	keyFunc, err := auth.NewCognitoKeyFunc(ctx, cfg.JWTIssuer)
	if err != nil {
		return nil, fmt.Errorf("build auth key resolver: %w", err)
	}
	verifier := auth.NewVerifier(cfg.JWTIssuer, cfg.JWTClientID, keyFunc)
	return middleware.NewAuthMiddleware(verifier).Wrap(routes), nil
}
