package lambdaclient

import (
	"errors"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/s3"
)

var LambdaClient *lambda.Lambda
var LambdaSession *session.Session

var dynamoClient *dynamodb.DynamoDB
var s3Client *s3.S3

func Initialize() (err error) {
	if os.Getenv("DYNAMOTABLE") == "" {
		err = errors.New("no DYNAMOTABLE environment variable")
		return
	}

	if os.Getenv("BUCKET") == "" {
		err = errors.New("no BUCKET environment variable")
		return
	}

	if os.Getenv("OBJECTPREFIX") == "" {
		err = errors.New("no OBJECTPREFIX environment variable")
		return
	}

	log.Printf("create session")
	sess, err := session.NewSession()
	if err != nil {
		return
	}

	LambdaSession = sess

	log.Printf("create Lambda client")
	LambdaClient = lambda.New(sess)

	log.Printf("create DynamoDB client")
	dynamoClient = dynamodb.New(LambdaSession)

	log.Printf("create S3 client")
	s3Client = s3.New(LambdaSession)

	return
}
