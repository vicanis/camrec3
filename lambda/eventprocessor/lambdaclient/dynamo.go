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

func GetUnprocessedEvents(limit int64) (items []Event, err error) {
	filt := expression.Name("processed").Equal(expression.Value(false))

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
		Limit:                     aws.Int64(limit),
	})
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
	}

	return
}

func (e Event) SetProcessed() (err error) {
	keyCond := expression.Key("id").Equal(expression.Value(e.Id))

	upd := expression.UpdateBuilder{}
	upd.Set(expression.Name("processed"), expression.Value(true))

	expr, err := expression.NewBuilder().WithKeyCondition(keyCond).WithUpdate(upd).Build()
	if err != nil {
		err = fmt.Errorf("expression build failed: %w", err)
		return
	}

	_, err = dynamoClient.UpdateItem(&dynamodb.UpdateItemInput{
		TableName:                 aws.String(os.Getenv("DYNAMOTABLE")),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
	})
	if err != nil {
		err = fmt.Errorf("update item failed: %w", err)
		return
	}

	return
}
