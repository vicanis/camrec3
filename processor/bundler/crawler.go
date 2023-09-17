package bundler

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

type chunkFile struct {
	name string
	ts   time.Time
}

type chunkFiles []chunkFile

const dirRaw = "/video/raw"

func getChunkFiles() (files chunkFiles, err error) {
	files = make(chunkFiles, 0)

	checkDirs := []string{}

	y, m, d := time.Now().Date()

	checkDirs = append(checkDirs, fmt.Sprintf(
		"%s/%04d/%02d/%02d",
		dirRaw, y, m, d,
	))

	y, m, d = time.Now().Add(-24 * time.Hour).Date()
	checkDirs = append(checkDirs, fmt.Sprintf(
		"%s/%04d/%02d/%02d",
		dirRaw, y, m, d,
	))

	var walk func(string)

	walk = func(dir string) {
		dirFiles, err := os.ReadDir(dir)
		if err != nil {
			return
		}

		for i, f := range dirFiles {
			fileType := f.Type()

			if fileType.IsDir() {
				path := dir + "/" + f.Name()

				ok := true

				for _, checkDir := range checkDirs {
					if checkDir == path {
						break
					}

					if !strings.HasPrefix(checkDir, path) {
						ok = false
						break
					}
				}

				if ok {
					walk(path)
				}

				continue
			}

			if !fileType.IsRegular() {
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
	}

	walk(dirRaw)

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
	y, m, d := f.ts.Date()
	return os.ReadFile(
		fmt.Sprintf(
			"%s/%04d/%02d/%02d/%s",
			dirRaw, y, m, d, f.name,
		),
	)
}
