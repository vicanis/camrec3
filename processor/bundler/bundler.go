package bundler

import (
	"errors"
	"fmt"
	"os"
	"processor/encoder"
	"time"
)

func SearchVideoBundle(ts time.Time) (data []byte, err error) {
	chunkStart := ts.Add(-3 * time.Second)

	data, err = fetchChunkData(chunkStart)
	if err != nil {
		return
	}

	_, _, s := chunkStart.Clock()

	data, err = encoder.Encode(data, s)
	if err != nil {
		return
	}

	y, m, d := ts.Date()
	dir := fmt.Sprintf("../video/chunks/%04d/%02d/%02d", y, m, d)

	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return
	}

	file := ts.Format("2006-01-02-15-04-05") + ".mp4"

	err = os.WriteFile(fmt.Sprintf(dir+"/"+file), data, 0644)

	return
}

func fetchChunkData(ts time.Time) (data []byte, err error) {
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
