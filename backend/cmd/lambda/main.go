// Lambda entrypoint behind an API Gateway HTTP API (payload format 2.0).
package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"

	"github.com/kazemisoroush/direct/backend/internal/api"
)

func main() {
	adapter := httpadapter.NewV2(api.New())
	lambda.Start(adapter.ProxyWithContext)
}
