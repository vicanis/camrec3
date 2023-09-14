package database

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

var sess *session.Session
var dynamoClient *dynamodb.DynamoDB

type ItemList struct {
	Count int   `json:"count"`
	Items Items `json:"items"`
}

type Items []Item

type Item struct {
	Id        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Processed bool      `json:"processed"`
	Raw       string    `json:"raw"`
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

func GetItems(day time.Time) (items ItemList, err error) {
	filt := expression.Name("unix").Between(
		expression.Value(day.Unix()),
		expression.Value(day.Add(24*time.Hour).Unix()),
	)
	expr, err := expression.NewBuilder().WithFilter(filt).Build()
	if err != nil {
		err = fmt.Errorf("expression build failed: %w", err)
		return
	}

	scan, err := dynamoClient.Scan(&dynamodb.ScanInput{
		TableName:                 aws.String(os.Getenv("DYNAMOTABLE")),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
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
			err = fmt.Errorf("database document unmarshal error: %w", err)
			return
		}

		items.Items = append(items.Items, item)
	}

	sort.Sort(items.Items)

	return
}

func (t Items) Len() int {
	return len(t)
}

func (t Items) Less(a, b int) bool {
	return t[a].Timestamp.After(t[b].Timestamp)
}

func (t Items) Swap(a, b int) {
	t[a], t[b] = t[b], t[a]
}
