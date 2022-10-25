package main

import (
	"context"
	"log"
	"mgmt/handlers"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var err error

var ginLambda *ginadapter.GinLambda

func LambdaHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}

func main() {
	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	if gin.Mode() == "debug" {
		r.Use(gin.Logger())
	}

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	err = r.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal(err)
	}

	mgmt := r.Group("/mgmt")
	{
		mgmt.POST("/orders", handlers.Testfunc2) // same as /mgmt/orders
		mgmt.POST("/buy", handlers.Testfunc2)    // same as /mgmt/buy
	}

	if os.Getenv("GO_ENV") == "local" {
		err = r.Run(":8080")
		if err != nil {
			log.Fatal("Function Run() has failed with error:", err)
		}
	} else {
		ginLambda = ginadapter.New(r)
		lambda.Start(LambdaHandler)
	}
}
