package main

import (
	"context"
	"lambda/lambdaclient"
	"lambda/parser"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	runtime "github.com/aws/aws-lambda-go/lambda"
)

func init() {
	bucketId := os.Getenv("BUCKET")
	if bucketId == "" {
		log.Fatal("no BUCKET environment variable")
	}

	filePrefix := os.Getenv("OBJECTPREFIX")
	if filePrefix == "" {
		log.Fatal("no OBJECTPREFIX environment variable")
	}
}

func main() {
	runtime.Start(handleRequest)
}

func handleRequest(ctx context.Context, event events.SimpleEmailEvent) (string, error) {
	err := processEvents(event)
	if err != nil {
		return "ERROR", err
	}

	return "OK", nil
}

func processEvents(event events.SimpleEmailEvent) error {
	for _, record := range event.Records {
		id := record.SES.Mail.MessageID

		log.Printf("process message %s", id)

		body, err := lambdaclient.GetMailData(id)
		if err != nil {
			return err
		}

		ts, err := parser.ParseTimestamp(body)
		if err != nil {
			return err
		}

		log.Printf("> timestamp parsed: %s", ts.Format(time.RFC1123))
	}

	return nil
}
