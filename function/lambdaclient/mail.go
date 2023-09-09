package lambdaclient

import (
	"fmt"
	"io"
	"log"
	"mime/quotedprintable"
	"net/mail"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

var s3Client *s3.S3

func init() {
	if os.Getenv("BUCKET") == "" {
		log.Fatal("no BUCKET environment variable")
	}

	if os.Getenv("OBJECTPREFIX") == "" {
		log.Fatal("no OBJECTPREFIX environment variable")
	}

	s3Client = s3.New(LambdaSession)
}

func GetMailData(id string) (data string, err error) {
	obj, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("BUCKET")),
		Key: aws.String(
			fmt.Sprintf(
				"%s/%s",
				os.Getenv("OBJECTPREFIX"),
				id,
			),
		),
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

	data = string(body)

	return
}
