package main

import (
	"context"
	"log"
	"mgmt/handlers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var err error

var ginLambda *ginadapter.GinLambda

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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

	// mgmt.Routes(r) // Add gin Engine to api/mgmt.go

	// mgmt := r.Group("/mgmt")
	// {
	// 	mgmt.POST("/orders", handlers.Testfunc2) // Default same as /forms
	// 	mgmt.POST("/buy", handlers.Testfunc2)    // same as /forms/test
	// }

	r.POST("/", handlers.Testfunc2) // same as /forms/test
	r.GET("/", handlers.Testfunc2) // same as /forms/test

	// err = r.Run(":8080")
	// if err != nil {
	// log.Fatal("Fucntion Run() has failed with error:", err)
	// }

	ginLambda = ginadapter.New(r)
	lambda.Start(Handler)
}
