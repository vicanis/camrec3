package main

import (
	"context"
	"errors"
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

	items, err := lambdaclient.GetUnprocessedEvents(1)
	if err != nil {
		return "ERROR", err
	}

	if len(items) == 0 {
		return "ERROR", errors.New("not found")
	}

	item := items[0]

	err = item.SetProcessed()
	if err != nil {
		return "ERROR", err
	}

	return "OK", nil
}
