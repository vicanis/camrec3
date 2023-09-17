package stream

import (
	"fmt"
	"os"
	"time"
)

func saveChunk(chunk []byte) (err error) {
	y, m, d := time.Now().Date()

	dir := fmt.Sprintf(
		"/video/raw/%04d/%02d/%02d",
		y, m, d,
	)

	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return
	}

	err = os.WriteFile(
		fmt.Sprintf(
			"%s/%s",
			dir,
			time.Now().Format("2006-01-02-15-04-05"),
		),
		chunk,
		0644,
	)
	if err != nil {
		return
	}

	return
}
