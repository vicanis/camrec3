package main

import (
	"context"
	"log"
	"mailhandler/lambdaclient"
	"mailhandler/parser"
	"time"

	"github.com/aws/aws-lambda-go/events"
	runtime "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	if err := lambdaclient.Initialize(); err != nil {
		log.Fatal(err)
	}

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
			UnixTime:  ts.Unix(),
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
