// Lambda entrypoint behind an API Gateway HTTP API (payload format 2.0).
package main

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"

	"github.com/kazemisoroush/direct/backend/internal/api"
	appconfig "github.com/kazemisoroush/direct/backend/internal/config"
	"github.com/kazemisoroush/direct/backend/internal/restaurant"
)

// startupTimeout bounds cold-start dependency I/O so a hung lookup fails fast instead of
// stalling the whole Lambda init window. Per-request work uses the invocation's own context.
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

	proxy := httpadapter.NewV2(handler)
	lambda.Start(proxy.ProxyWithContext)
}
