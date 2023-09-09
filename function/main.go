package main

import (
	"context"
	"lambda/lambdaclient"
	"lambda/parser"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
	runtime "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambdaclient.Initialize()
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

		evt := lambdaclient.Event{
			Id:        id,
			Timestamp: *ts,
			Raw:       body,
		}

		log.Printf("> save event data")

		err = lambdaclient.SaveEvent(evt)
		if err != nil {
			return err
		}

		log.Printf("> ok")
	}

	return nil
}
