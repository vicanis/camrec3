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

	items, err := lambdaclient.GetUnprocessedEvents(50)
	if err != nil {
		return "ERROR", err
	}

	if len(items) == 0 {
		return "ERROR", errors.New("not found")
	}

	for _, item := range items {
		err := lambdaclient.ProcessEvent(item)
		if err != nil {
			return "ERROR", err
		}
	}

	return "OK", nil
}
