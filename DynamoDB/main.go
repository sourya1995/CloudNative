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
	CreateTableLsi(TableName, svc)
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

		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("id-sdk"),
				KeyType: aws.String("HASH"),
			},
			{
				AttributeName: aws.String("name"),
				KeyType: aws.String("RANGE"),
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

func CreateRecord(tableName string, svc dynamodbiface.DynamoDBAPI){
	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]*dynamodb.AttributeValue{
			"partitionkey": {
				S: aws.String("test"),
			}
			"sortkey": {
				S: aws.String("2022"),
			},

		},
	}

	_, err := svc.PutItem(input)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Result created successfully")
}

func QueryRecord(tableName string, svc dynamodbiface.DynamoDBAPI){
	keyCond := expression.KeyAnd(
		expression.Key("partitionkey").Equal(expression.Value("test")),
		expression.Key("sortkey").BeginsWith("2"),
	)

	expr, err := expression.NewBuilder().
		WithKeyCondition(keyCond).
		Build()

	query := &dynamodb.QueryInput {
		ExpressionAttributeNames: expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression: expr.KeyCondition(),
		TableName: aws.String(tableName),
	}

	output, err := svc.Query(query)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Result querying record: ")
	for _, v := range output.Items {
		fmt.Printf("value: %s \n", v)
	}
}

func CreateTableLsi(tableName string, svc dynamodbiface.DynamoDBAPI) {
	fmt.Println("Creating table LSI...")
	billingMode := "PAY_PER_REQUEST"

	params := &dynamodb.CreateTableInput{
		TableName: aws.String(tableName),

		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("id-sdk"),
				AttributeType: aws.String("S"),

			},

			{
				AttributeName: aws.String("name"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("date"),
				AttributeType: aws.String("S"),
			},
		},

		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("id-sdk"),
				KeyType: aws.String("HASH"),
			},
			{
				AttributeName: aws.String("name"),
				keyType: aws.String("RANGE"),
			},
		},

		LocalSecondaryIndexes: []*dynamodb.LocalSecondaryIndex {
			{
				IndexName: aws.String("my-index-from-sdk"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("id-sdk"),
						KeyType: aws.String("HASH"),
					},
					{
						AttributeName: aws.String("date"),
						KeyType: aws.String("RANGE"),
					},
				},

				Projection: &dynamodb.Projection{
					ProjectionType: aws.String("KEYS_ONLY"),

				},

			},
		},
		BillingMode: &billingMode,

	}

	output, err := svc.CreateTable(params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Result:")
	fmt.Println(output)
	fmt.Println("LSI created successfully")
}

func CreateTableGsi(tableName string, svc dynamodbiface.DynamoDBAPI){
	fmt.Println("Creating GSI...")
	billingMode := "PROVISIONED"

	params := &dynamodb.CreateTableInput {
		TableName: aws.String(tableName),

		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("id-sdk"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("name"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("date"),
				AttributeType: aws.String("S"),
			},
		},

		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("id-sdk"),
				KeyType: aws.String("HASH"),
			},
			{
				AttributeName: aws.String("name"),
				KeyType: aws.String("RANGE"),
			},

		},

		GlobalSecondaryIndexes: []* dynamodb.GlobalSecondaryIndex{
			{
				IndexName: aws.String("gsi-from-sdk"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("id-sdk"),
						KeyType: aws.String("HASH"),
					},
					{
						AttributeName: aws.String("date"),
						KeyType: aws.String("RANGE"),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String("ALL"),
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput {
					ReadCapacityUnits: aws.Int64(1),
					WriteCapacityUnits: aws.Int64(1),
				},

			},
		},

		BillingMode: &billingMode,
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits: aws.Int64(10),
			WriteCapacityUnits: aws.Int64(5),
		}



	}

	output, err := svc.CreateTable(params)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Result: ")
	fmt.Println(output)

	fmt.Println("finished")
}