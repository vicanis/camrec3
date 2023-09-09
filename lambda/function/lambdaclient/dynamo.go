package lambdaclient

import (
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Event struct {
	Id        string    `json:"id"`
	Raw       string    `json:"raw"`
	Timestamp time.Time `json:"timestamp"`
	Processed bool      `json:"processed"`
}

func SaveEvent(evt Event) (err error) {
	av, err := dynamodbattribute.MarshalMap(evt)
	if err != nil {
		return
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(os.Getenv("DYNAMOTABLE")),
	}

	_, err = dynamoClient.PutItem(input)
	if err != nil {
		return
	}

	return
}
