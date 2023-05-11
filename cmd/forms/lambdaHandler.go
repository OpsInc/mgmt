package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

//nolint:wrapcheck
func lambdaHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}
