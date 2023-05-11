package main

import (
	"log"
	"mgmt/internal/configs"
	"mgmt/internal/database"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type application struct {
	db *dynamodb.Client
}

func main() {
	err := configs.GetActiveProfile()
	if err != nil {
		log.Fatal(err)
	}

	// We initialize all configs for our application
	app := &application{
		db: database.AWSConnection(), // Initialize DB connection
	}

	// Starts the gin server and routes configurations
	app.routes()
}
