// Package provides a microservice for forms processing.
// It can be used as a Lambda, Docker or locally
package main

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"mgmt/internal/configs"
	"mgmt/internal/database"
)

// Application settings
type application struct {
	db *dynamodb.Client
}

// Fetch current profile to load appropriate ENVs
// Initialize all configs for our application
// Start the gin server and routes configurations
func main() {
	err := configs.GetActiveProfile()
	if err != nil {
		log.Fatal(err)
	}

	app := &application{
		db: database.AWSConnection(), // Initialize DB connection
	}

	app.routes()
}
