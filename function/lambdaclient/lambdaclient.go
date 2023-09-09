package lambdaclient

import (
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

func Initialize() {
	if os.Getenv("DYNAMOTABLE") == "" {
		log.Fatal("no DYNAMOTABLE environment variable")
	}

	if os.Getenv("BUCKET") == "" {
		log.Fatal("no BUCKET environment variable")
	}

	if os.Getenv("OBJECTPREFIX") == "" {
		log.Fatal("no OBJECTPREFIX environment variable")
	}

	log.Printf("create session")
	sess, err := session.NewSession()
	if err != nil {
		log.Fatal(err)
	}

	LambdaSession = sess

	log.Printf("create Lambda client")
	LambdaClient = lambda.New(sess)

	log.Printf("create DynamoDB client")
	dynamoClient = dynamodb.New(LambdaSession)

	log.Printf("create S3 client")
	s3Client = s3.New(LambdaSession)
}
