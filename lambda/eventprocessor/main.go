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

	return "OK", nil
}
