package lambdaclient

import (
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

var LambdaClient *lambda.Lambda
var LambdaSession *session.Session

func init() {
	sess, err := session.NewSession()
	if err != nil {
		log.Fatal(err)
	}

	LambdaSession = sess
	LambdaClient = lambda.New(sess)
}
