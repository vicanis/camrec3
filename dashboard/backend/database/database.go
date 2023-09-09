package database

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var sess *session.Session
var dynamoClient *dynamodb.DynamoDB

type Item struct {
	Id        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Processed bool      `json:"processed"`
	Raw       string    `json:"raw"`
}

type ItemList struct {
	Count int    `json:"count"`
	Items []Item `json:"items"`
}

func Initialize() (err error) {
	sess, err = session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("REGION")),
	})
	if err != nil {
		err = fmt.Errorf("create session error: %w", err)
		return
	}

	dynamoClient = dynamodb.New(sess)

	return
}

func GetItems() (items ItemList, err error) {
	scan, err := dynamoClient.Scan(&dynamodb.ScanInput{
		TableName: aws.String(os.Getenv("DYNAMOTABLE")),
	})
	if err != nil {
		err = fmt.Errorf("database scan failed: %w", err)
		return
	}

	items = ItemList{
		Count: int(*scan.Count),
		Items: make([]Item, 0),
	}

	for _, rawItem := range scan.Items {
		item := Item{}

		err = dynamodbattribute.UnmarshalMap(rawItem, &item)
		if err != nil {
			err = fmt.Errorf("database document unmarshal error: %s", err)
			return
		}

		items.Items = append(items.Items, item)
	}

	return
}
