// Local HTTP server for running the Direct API outside Lambda during development.
package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/kazemisoroush/direct/backend/internal/api"
	appconfig "github.com/kazemisoroush/direct/backend/internal/config"
	"github.com/kazemisoroush/direct/backend/internal/restaurant"
)

// startupTimeout bounds cold-start dependency I/O so a hung lookup fails fast.
const startupTimeout = 15 * time.Second

func main() {
	cfg := appconfig.Load()

	startupCtx, cancel := context.WithTimeout(context.Background(), startupTimeout)
	defer cancel()

	awsCfg, err := config.LoadDefaultConfig(startupCtx)
	if err != nil {
		log.Fatalf("load AWS config: %v", err)
	}

	store := restaurant.NewDynamoStore(dynamodb.NewFromConfig(awsCfg), cfg.Table)

	// api.New builds the Cognito JWKS resolver, whose context governs the lifetime of the
	// hourly key refresh — it must outlive startup, so use Background here, not startupCtx.
	handler, err := api.New(context.Background(), cfg, store)
	if err != nil {
		log.Fatalf("configure api: %v", err)
	}

	log.Printf("Direct API listening on %s", cfg.ServerAddr())
	log.Fatal(http.ListenAndServe(cfg.ServerAddr(), handler))
}
