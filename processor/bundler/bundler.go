package bundler

import (
	"errors"
	"time"
)

func SearchVideoBundle(ts time.Time) (data []byte, err error) {
	chunkFiles, err := getChunkFiles()
	if err != nil {
		return
	}

	if len(chunkFiles) == 0 {
		err = errors.New("no chunk files")
		return
	}

	if chunkFiles[0].ts.After(ts) || chunkFiles[len(chunkFiles)-2].ts.Before(ts) {
		err = errors.New("timestamp is not in the buffer range")
		return
	}

	data = make([]byte, 0)

	for i, chunk := range chunkFiles {
		if chunk.ts.After(ts) {
			for j := i; j <= i+1; j++ {
				buf, err := chunkFiles[j].Read()
				if err != nil {
					break
				}

				data = append(data, buf...)
			}

			break
		}
	}

	return
}
