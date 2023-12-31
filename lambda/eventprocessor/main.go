package main

import (
	"context"
	"eventprocessor/lambdaclient"
	"log"

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
	log.Printf("start")

	items, err := lambdaclient.GetUnprocessedEvents(10)
	if err != nil {
		return "ERROR", err
	}

	if len(items) == 0 {
		return "OK", nil
	}

	for _, item := range items {
		err := lambdaclient.ProcessEvent(item)
		if err != nil {
			log.Printf("event processing failed: %s", err)
		}
	}

	return "OK", nil
}
