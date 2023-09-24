package main

import (
	"context"
	"dashboard/lambdaclient"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	runtime "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	if err := lambdaclient.Initialize(); err != nil {
		log.Fatal(err)
	}

	runtime.Start(wrapper(handleRequest))
}

type SimpleRequestHandler func(event events.APIGatewayProxyRequest) (any, error)
type AwsRequestHandler func(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

func wrapper(handler SimpleRequestHandler) AwsRequestHandler {
	return func(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		data, err := handler(event)
		if err != nil {
			buf, _ := json.Marshal(ResponseError{
				Error: err.Error(),
			})

			return events.APIGatewayProxyResponse{
				Body:       string(buf),
				StatusCode: http.StatusInternalServerError,
			}, nil
		}

		if binary, ok := data.([]byte); ok {
			return events.APIGatewayProxyResponse{
				Body:       string(binary),
				StatusCode: http.StatusOK,
			}, nil
		}

		buf, err := json.Marshal(data)
		if err != nil {
			buf, _ = json.Marshal(ResponseError{
				Error: err.Error(),
			})

			return events.APIGatewayProxyResponse{
				Body:       string(buf),
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		return events.APIGatewayProxyResponse{
			Body:       string(buf),
			StatusCode: http.StatusOK,
		}, nil
	}
}

func handleRequest(event events.APIGatewayProxyRequest) (any, error) {
	action := event.QueryStringParameters["action"]
	if action == "" {
		return nil, errors.New("no action")
	}

	if event.HTTPMethod == http.MethodGet {
		if action == "list" {
			return handleEventList(event)
		}

		if action == "load" {
			return handleEventData(event)
		}
	}

	return nil, errors.New("bad action")
}

func handleEventList(event events.APIGatewayProxyRequest) (list *ResponseList, err error) {
	qsa := event.QueryStringParameters

	date := qsa["date"]
	if date == "" {
		err = errors.New("no date")
		return
	}

	ts, err := time.Parse("20060102150405", date)
	if err != nil {
		return
	}

	events, err := lambdaclient.GetEventList(ts)
	if err != nil {
		return
	}

	list = &ResponseList{
		Date:   ts,
		Events: events,
	}

	return
}

func handleEventData(event events.APIGatewayProxyRequest) (data []byte, err error) {
	qsa := event.QueryStringParameters

	file := qsa["file"]
	if file == "" {
		err = errors.New("no file")
		return
	}

	return lambdaclient.LoadFile(file)
}

type ResponseError struct {
	Error string `json:"error"`
}

type ResponseList struct {
	Date   time.Time            `json:"date"`
	Events []lambdaclient.Event `json:"events"`
}
