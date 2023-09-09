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

func GetMailData(id string) (data string, err error) {
	s3client := s3.New(LambdaSession)

	obj, err := s3client.GetObject(&s3.GetObjectInput{
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
