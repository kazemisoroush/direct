// CDK app defining the Direct walking-skeleton stack.
//
// M0 is the baseline: a static frontend served over CloudFront and a single health
// Lambda behind an HTTP API. Auth (Cognito), restaurants and ordering land in later
// milestones — this stack stays deliberately empty of product resources.
package main

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations"
	golambda "github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdklabs/cdk-nag-go/cdknag/v2"
)

// NewDirectStack defines the frontend hosting and the health API for the M0 baseline.
func NewDirectStack(scope constructs.Construct, id string, props *awscdk.StackProps) awscdk.Stack {
	stack := awscdk.NewStack(scope, &id, props)

	webOrigin := stack.Node().TryGetContext(jsii.String("webOrigin"))
	origin, ok := webOrigin.(string)
	if !ok || origin == "" {
		origin = "http://localhost:3000"
	}

	// The frontend hosting is created first so its CloudFront origin can be allowed by
	// the API CORS below. localhost stays allowed for local dev.
	hosting := newFrontendHosting(stack)
	allowedOrigins := jsii.Strings(origin, hosting.URL())

	fn := golambda.NewGoFunction(stack, jsii.String("Api"), &golambda.GoFunctionProps{
		Entry:   jsii.String("../backend/cmd/lambda"),
		Timeout: awscdk.Duration_Seconds(jsii.Number(30)),
	})

	api := awsapigatewayv2.NewHttpApi(stack, jsii.String("HttpApi"), &awsapigatewayv2.HttpApiProps{
		CorsPreflight: &awsapigatewayv2.CorsPreflightOptions{
			AllowOrigins: allowedOrigins,
			AllowMethods: &[]awsapigatewayv2.CorsHttpMethod{
				awsapigatewayv2.CorsHttpMethod_GET,
				awsapigatewayv2.CorsHttpMethod_POST,
				awsapigatewayv2.CorsHttpMethod_PATCH,
				awsapigatewayv2.CorsHttpMethod_DELETE,
			},
			AllowHeaders: jsii.Strings("Content-Type", "Authorization"),
		},
	})

	integration := awsapigatewayv2integrations.NewHttpLambdaIntegration(jsii.String("ApiIntegration"), fn, nil)
	healthRoutes := api.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Path:        jsii.String("/health"),
		Methods:     &[]awsapigatewayv2.HttpMethod{awsapigatewayv2.HttpMethod_GET},
		Integration: integration,
	})

	// Upload the built site plus a config.json rendered from the stack outputs, so the SPA
	// reads its API settings at runtime and never drifts from the backend.
	hosting.deploy(stack, api.Url())

	awscdk.NewCfnOutput(stack, jsii.String("FrontendUrl"), &awscdk.CfnOutputProps{Value: jsii.String(hosting.URL())})
	awscdk.NewCfnOutput(stack, jsii.String("ApiUrl"), &awscdk.CfnOutputProps{Value: api.Url()})

	suppressNag(stack, (*healthRoutes)[0])

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)
	NewDirectStack(app, "DirectStack", nil)
	NewDirectCICDStack(app, "DirectCICDStack", &awscdk.StackProps{
		Env: &awscdk.Environment{
			Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
			Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
		},
	})
	awscdk.Aspects_Of(app).Add(cdknag.NewAwsSolutionsChecks(&cdknag.NagPackProps{Verbose: jsii.Bool(true)}), nil)
	app.Synth(nil)
}
