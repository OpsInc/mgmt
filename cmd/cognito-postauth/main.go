package main

import (
	"fmt"
	"mgmt/internal/database"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	ID    string
	Email string
}

// map[callerContext:map[awsSdkVersion:aws-sdk-unknown-unknown clientId:6ifqsiq3kud2pcjrv5csiga7oc] region:us-east-1 request:map[userAttributes:map[email:test@gmail.com] validationData:<nil>] response:map[autoConfirmUser:false autoVerifyEmail:false autoVerifyPhone:false] triggerSource:PreSignUp_SignUp userName:5 userPoolId:us-east-1_VaxI7rfIr version:1]
func handler(event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	fmt.Printf("PostConfirmation for user: %s\n", event.UserName)

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
