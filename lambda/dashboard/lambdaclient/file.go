package lambdaclient

import (
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func LoadFile(file string) (data []byte, err error) {
	obj, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("BUCKET")),
		Key:    aws.String(file),
	})
	if err != nil {
		return
	}

	return io.ReadAll(obj.Body)
}
