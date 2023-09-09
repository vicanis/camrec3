package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/quotedprintable"
	"net/mail"
	"os"
	"regexp"
	"time"

	"github.com/aws/aws-lambda-go/events"
	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/s3"
)

var awsClient *lambda.Lambda
var awsSession *session.Session

func init() {
	bucketId := os.Getenv("BUCKET")
	if bucketId == "" {
		log.Fatal("no BUCKET environment variable")
	}

	filePrefix := os.Getenv("OBJECTPREFIX")
	if filePrefix == "" {
		log.Fatal("no OBJECTPREFIX environment variable")
	}

	sess, err := session.NewSession()
	if err != nil {
		log.Fatal(err)
	}

	awsSession = sess
	awsClient = lambda.New(sess)
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

func processEvents(event events.SimpleEmailEvent) (err error) {
	for _, record := range event.Records {
		if err = processMessage(record); err != nil {
			return
		}
	}

	return
}

func processMessage(record events.SimpleEmailRecord) (err error) {
	messageId := record.SES.Mail.MessageID

	log.Printf("message ID: %s", messageId)

	s3client := s3.New(awsSession)

	obj, err := s3client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("BUCKET")),
		Key:    aws.String(fmt.Sprintf("%s/%s", os.Getenv("OBJECTPREFIX"), messageId)),
	})
	if err != nil {
		return
	}

	parsed, err := mail.ReadMessage(obj.Body)
	if err != nil {
		return
	}

	decoder := quotedprintable.NewReader(parsed.Body)

	body, err := io.ReadAll(decoder)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("decoded body: %s", body)

	ts := ParseTimestamp(string(body))

	log.Printf("> timestamp parsed: %s", ts.Format(time.RFC1123))

	return
}

func ParseTimestamp(body string) (ts time.Time) {
	rx := regexp.MustCompile(`(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2})`)

	matches := rx.FindAllString(string(body), 1)
	if len(matches) == 0 {
		log.Fatal("no timestamp")
	}

	ts, err := time.ParseInLocation("2006-01-02 15:04:05", matches[0], time.Local)
	if err != nil {
		log.Fatal(err)
	}

	return ts
}
