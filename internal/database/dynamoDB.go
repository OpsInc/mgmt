package database

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
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

// func PutItems(db *dynamodb.Client, items map[string]types.AttributeValue) {

// - Takes a struct data in order to convert it as map[string]types.AttributeValue items
// - Writes the items in DynamoDB table
// - The DynamoDB table is defined by env var "DATABASE_NAME".
func PutItems(db *dynamodb.Client, data any) {
	databaseName := os.Getenv("DATABASE_NAME")

	if databaseName == "" {
		log.Fatal("No database has been set as env variable: DATABASE_NAME")
	}

	// items := &dynamodb.PutItemInput{
	// 	TableName: aws.String(database_name),
	// 	Item: map[string]types.AttributeValue{
	// 		"id":        &types.AttributeValueMemberS{Value: "3"},
	// 		"name":      &types.AttributeValueMemberS{Value: f.Name},
	// 		"age":       &types.AttributeValueMemberS{Value: f.Age},
	// 		"email":     &types.AttributeValueMemberS{Value: f.Email},
	// 		"frequency": &types.AttributeValueMemberS{Value: f.Frequency},
	// 		"goal":      &types.AttributeValueMemberS{Value: f.Goal},
	// 	},
	// }

	items, err := attributevalue.MarshalMap(data)
	if err != nil {
		log.Fatal("Empty dynamodb items provided with error:", err)
	}

	out, err := db.PutItem(context.TODO(), &dynamodb.PutItemInput{ //nolint:exhaustruct
		TableName: aws.String(databaseName),
		Item:      items,
	})
	if err != nil {
		log.Fatal("Database:", databaseName, "has PutItems Error: ", err)
	}

	log.Println(out.Attributes)
}

// func ListTables(cfg aws.Config) {
//
//   conn := dynamodb.NewFromConfig(cfg)
//
//   log.Println(conn.ListTables(context.Context))
// }
