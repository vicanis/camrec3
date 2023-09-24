package lambdaclient

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type Event struct {
	Id        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Processed bool      `json:"processed"`
	File      string    `json:"file"`
}

type EventList []Event

func GetEventList(ts time.Time) (list EventList, err error) {
	y, m, d := ts.Date()

	filt := expression.And(
		expression.Name("unix").GreaterThan(
			expression.Value(
				time.Date(y, m, d, 0, 0, 0, 0, time.Local).Unix(),
			),
		),
		expression.Name("unix").LessThan(
			expression.Value(
				time.Date(y, m, d, 23, 59, 59, 0, time.Local).Unix(),
			),
		),
	)

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

	list = make(EventList, 0)

	for _, rawItem := range scan.Items {
		item := Event{}

		err = dynamodbattribute.UnmarshalMap(rawItem, &item)
		if err != nil {
			err = fmt.Errorf("database document unmarshal error: %w", err)
			return
		}

		list = append(list, item)
	}

	sort.Sort(list)

	return
}

func (e EventList) Len() int {
	return len(e)
}

func (e EventList) Less(a, b int) bool {
	return e[a].Timestamp.After(e[b].Timestamp)
}

func (e EventList) Swap(a, b int) {
	e[a], e[b] = e[b], e[a]
}
