package database

import (
	"context"
	"log"
	"mgmt/views"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func AWSConnection() *dynamodb.Client {
	var (
		cfg aws.Config
		err error
	)

	awsProfile := os.Getenv("AWS_PROFILE")

	if awsProfile == "" {
		cfg, err = config.LoadDefaultConfig(context.TODO())
		if err != nil {
			log.Fatal("Unable to load AWS profile from LAMBDA:", err)
		}
	} else {
		cfg, err = config.LoadDefaultConfig(context.TODO(),
			config.WithSharedConfigProfile(awsProfile))
		if err != nil {
			log.Fatal("Unable to load AWS profile with error:", err)
		}
	}

	return dynamodb.NewFromConfig(cfg)
}

func ListTables(db *dynamodb.Client) {
	// input := &dynamodb.ListTablesInput{}   // Can be replaced by nil since we do not provide any config other than an empty struct

	tables, err := db.ListTables(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(tables.TableNames)

	for _, tables := range tables.TableNames {
		log.Println("Table name", tables)
	}
}

func PutItems(db *dynamodb.Client, f *views.FormOutput) {
	input := &dynamodb.PutItemInput{
		TableName: aws.String("mgmt"),
		Item: map[string]types.AttributeValue{
			"id":    &types.AttributeValueMemberS{Value: "7"},
			"fname": &types.AttributeValueMemberS{Value: f.FirstName},
			"lname": &types.AttributeValueMemberS{Value: f.LastName},
		},
	}

	out, err := db.PutItem(context.TODO(), input)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(out.Attributes)
}

// func ListTables(cfg aws.Config) {
//
//   conn := dynamodb.NewFromConfig(cfg)
//
//   log.Println(conn.ListTables(context.Context))
// }
