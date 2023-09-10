package bundler

import (
	"log"
	"os"
	"sort"
	"time"
)

type chunkFile struct {
	name string
	ts   time.Time
}

type chunkFiles []chunkFile

func getChunkFiles() (files chunkFiles, err error) {
	files = make(chunkFiles, 0)

	dirFiles, err := os.ReadDir("../video/raw")
	if err != nil {
		return
	}

	for i, f := range dirFiles {
		if !f.Type().IsRegular() {
			continue
		}

		name := f.Name()

		ts, err := time.ParseInLocation("2006-01-02-15-04-05", name, time.Local)
		if err != nil {
			log.Printf("file#%d: %s parse failed: %s", i, name, err)
			continue
		}

		files = append(files, chunkFile{
			name: name,
			ts:   ts,
		})
	}

	sort.Sort(files)

	return
}

func (f chunkFiles) Len() int {
	return len(f)
}

func (f chunkFiles) Less(a, b int) bool {
	return f[a].ts.Before(f[b].ts)
}

func (f chunkFiles) Swap(a, b int) {
	f[a], f[b] = f[b], f[a]
}

func (f chunkFile) Read() (data []byte, err error) {
	return os.ReadFile("../video/raw/" + f.name)
}
