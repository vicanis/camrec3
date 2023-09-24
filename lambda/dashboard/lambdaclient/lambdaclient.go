package lambdaclient

import (
	"errors"
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

	sess, err := session.NewSession()
	if err != nil {
		return
	}

	LambdaSession = sess

	LambdaClient = lambda.New(sess)
	dynamoClient = dynamodb.New(LambdaSession)
	s3Client = s3.New(LambdaSession)

	return
}
