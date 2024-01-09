package lambdaclient

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func ProcessEvent(e Event) (err error) {
	var fileName string

	processorServer := os.Getenv("PROCESSORSERVER")
	if processorServer == "" {
		err = errors.New("no PROCESSORSERVER environment variable")
		return
	}

	okProcessed := false

	defer func() {
		if !okProcessed {
			log.Printf("item processing failed, do not update flag")
			return
		}

		if err := e.SetProcessed(fileName); err != nil {
			log.Printf("update item processed failed: %s", err)
		}
	}()

	url := processorServer + e.Timestamp.Format("20060102150405")

	log.Printf("get event video by URL: %s", url)

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("unsuccessful status code: %d", resp.StatusCode)
		return
	}

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	fileName = fmt.Sprintf(
		"%s/%s.mp4",
		os.Getenv("EVENTPREFIX"),
		e.Timestamp.Format("2006-01-02 15:04:05"),
	)

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("BUCKET")),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(buf),
	})
	if err != nil {
		return
	}

	log.Printf("event video was saved")

	okProcessed = true

	return
}
