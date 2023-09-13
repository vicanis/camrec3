package parser

import (
	"errors"
	"regexp"
	"time"
)

func ParseTimestamp(body string) (ts *time.Time, err error) {
	rx := regexp.MustCompile(`(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2})`)

	matches := rx.FindAllString(string(body), 1)
	if len(matches) == 0 {
		err = errors.New("no timestamp")
		return
	}

	parsed, err := time.ParseInLocation("2006-01-02 15:04:05", matches[0], time.Local)
	if err != nil {
		return
	}

	ts = &parsed

	return
}
