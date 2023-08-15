package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"mgmt/internal/database"
)

type Response struct {
	ID    string
	Email string
}

func handler(event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	log.Printf("PostConfirmation for user: %s\n", event.UserName)

	response := new(Response)

	response.ID = event.UserName
	response.Email = event.Request.UserAttributes["email"]

	db := database.AWSConnection()

	database.PutItems(db, response)

	return event, nil
}

func main() {
	lambda.Start(handler)
}
