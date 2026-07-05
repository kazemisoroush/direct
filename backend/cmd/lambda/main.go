// Lambda entrypoint behind an API Gateway HTTP API (payload format 2.0).
package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"

	"github.com/kazemisoroush/direct/backend/internal/api"
	appconfig "github.com/kazemisoroush/direct/backend/internal/config"
	"github.com/kazemisoroush/direct/backend/internal/restaurant"
)

func main() {
	ctx := context.Background()
	cfg := appconfig.Load()

	awsCfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("load AWS config: %v", err)
	}

	store := restaurant.NewDynamoStore(dynamodb.NewFromConfig(awsCfg), cfg.Table)

	handler, err := api.New(ctx, cfg, store)
	if err != nil {
		log.Fatalf("configure api: %v", err)
	}

	proxy := httpadapter.NewV2(handler)
	lambda.Start(proxy.ProxyWithContext)
}
