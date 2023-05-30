package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	TableName     = "your-table-name"
	GlobalRegions = "us-west-2,us-east-1" // Add the regions where you want to create replicas
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"), // Specify the region where you want to create the global table
	})
	if err != nil {
		log.Fatal(err)
	}

	svc := dynamodb.New(sess)

	// Create global table input
	input := &dynamodb.CreateGlobalTableInput{
		GlobalTableName: aws.String(TableName),
		ReplicationGroup: []*dynamodb.Replica{
			{
				RegionName: aws.String("us-west-2"), // Specify the region for the first replica
			},
			{
				RegionName: aws.String("us-east-1"), // Specify the region for the second replica
			},
		},
	}

	// Create the global table
	_, err = svc.CreateGlobalTable(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Global table creation initiated.")
}
