package lambdaclient

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type Event struct {
	Id        string    `json:"id"`
	Raw       string    `json:"raw"`
	Timestamp time.Time `json:"timestamp"`
	UnixTime  int64     `json:"unix"`
	Processed bool      `json:"processed"`
}

func GetUnprocessedEvents(limit int) (items []Event, err error) {
	filt := expression.Name("processed").Equal(expression.Value(false))

	expr, err := expression.NewBuilder().WithFilter(filt).Build()
	if err != nil {
		err = fmt.Errorf("expression build failed: %w", err)
		return
	}

	input := &dynamodb.ScanInput{
		TableName:                 aws.String(os.Getenv("DYNAMOTABLE")),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	}

	scan, err := dynamoClient.Scan(input)
	if err != nil {
		err = fmt.Errorf("database scan failed: %w", err)
		return
	}

	items = make([]Event, 0)

	for _, rawItem := range scan.Items {
		item := Event{}

		err = dynamodbattribute.UnmarshalMap(rawItem, &item)
		if err != nil {
			err = fmt.Errorf("database document unmarshal error: %w", err)
			return
		}

		items = append(items, item)

		if len(items) == limit {
			break
		}
	}

	return
}

func (e Event) SetProcessed(fileName string) (err error) {
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(os.Getenv("DYNAMOTABLE")),
		ExpressionAttributeNames: map[string]*string{
			"#processed": aws.String("processed"),
			"#file":      aws.String("file"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":yes": {
				BOOL: aws.Bool(true),
			},
			":path": {
				S: aws.String(fileName),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(e.Id),
			},
			"timestamp": {
				S: aws.String(e.Timestamp.Format(time.RFC3339)),
			},
		},
		UpdateExpression: aws.String("SET #processed = :yes, #file = :path"),
	}

	_, err = dynamoClient.UpdateItem(input)
	if err != nil {
		err = fmt.Errorf("update item failed: %w", err)
		return
	}

	return
}
