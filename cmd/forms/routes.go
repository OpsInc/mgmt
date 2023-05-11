package main

import (
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var (
	err       error
	ginLambda *ginadapter.GinLambda
)

func (app *application) routes() {
	if os.Getenv("GIN_MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}

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

	// Disable proxie checks
	err = r.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal(err)
	}

	// Routes
	forms := r.Group("/forms")
	{
		forms.POST("/users", app.createUsers)
	}

	// Run localy or for Lambda
	if os.Getenv("GO_ENV") == "local" {
		err = r.Run(":8080")
		if err != nil {
			log.Fatal("Function Run() has failed with error:", err)
		}
	} else {
		ginLambda = ginadapter.New(r)
		lambda.Start(lambdaHandler)
	}
}
