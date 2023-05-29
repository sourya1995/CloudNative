package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	_ "github.com/godror/godror"
)

const (
	OracleDBConnectionString = "user/password@host:port/service_name"
	Region                    = "us-west-2"
	TableName                 = "your-dynamodb-table"
)

func main() {
	// Establish a connection to OracleDB
	oracleDB, err := sql.Open("godror", OracleDBConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer oracleDB.Close()

	// Create a DynamoDB client
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(Region),
	})
	if err != nil {
		log.Fatal(err)
	}
	dynamoDB := dynamodb.New(sess)

	// Query data from OracleDB
	rows, err := oracleDB.Query("SELECT * FROM your_table")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterate over the rows and insert into DynamoDB
	for rows.Next() {
		var (
			id    string
			name  string
			email string
		)

		err := rows.Scan(&id, &name, &email)
		if err != nil {
			log.Println(err)
			continue
		}

		// Create a DynamoDB PutItem input
		input := &dynamodb.PutItemInput{
			TableName: aws.String(TableName),
			Item: map[string]*dynamodb.AttributeValue{
				"id":    {S: aws.String(id)},
				"name":  {S: aws.String(name)},
				"email": {S: aws.String(email)},
			},
		}

		// Insert item into DynamoDB
		_, err = dynamoDB.PutItem(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				log.Println(aerr.Error())
			} else {
				log.Println(err.Error())
			}
		}
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data migration complete.")
}
