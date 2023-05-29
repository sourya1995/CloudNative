package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

const (	 //replace with creds
	Region = "{{Region}}"
	AccessKey = "{{AccessKey}}"
	SecretKey = "{{SecretKey}}"
	Token     = ""
	TableName = "table-keys-work"
)


func main() {
	svc := CreateSession()
	CreateTableOnDemand(TableName, svc)
	CreateRecord(TableName, svc)
	QueryRecord(TableName, svc)
}

func CreateSession() dynamodbiface.DynamoDBAPI {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(Region),
		Credentials: credentials.NewStaticCredentials(AccessKey, SecretKey, Token),
	})

	if err != nil {
		fmt.Println(err)
	}	

	var svc dynamodbiface.DynamoDBAPI
	svc = dynamodb.New(sess)

	return svc
}

func CreateTableOnDemand(tableName string, svc dynamodbiface.DynamoDB) {
	fmt.Println("Create Table starts")
	billingMode := "PAY_PER_REQUEST"

	params := &dynamodb.CreateTableInput {
		TableName: aws.String(tableName),

		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{ 
			AttributeName: aws.String("partitionkey"),
			AttributeType: aws.String("S")
		},
		{
			AttributeName: aws.String("sortkey"),
			AttributeType: aws.String("S")
		},
		},
		BillingMode := &billingMode,
	}
	output, err := svc.CreateTable(params)
	if err != nil {
		fmt.Println(err)
	}

	description := dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	}
	svc.WaitUntilTableExists(&description)
	fmt.Printf("CREATING:\n Table ARN: %s \n, Table name: %s \n",*output.TableDescription.TableArn, *output.TableDescription.TableName)
}