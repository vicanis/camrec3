package stream

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

var cmd *exec.Cmd
var lastDataTimestamp time.Time = time.Now()

func Start(ctx context.Context) chan error {
	done := make(chan error, 1)

	go func() {
		done <- startStreaming(ctx)
	}()

	go func() {
		time.Sleep(time.Minute)

		timer := time.NewTicker(15 * time.Second)
		defer timer.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-timer.C:
				if err := checkProcessState(); err != nil {
					log.Print(err)
					done <- err
					return
				}
			}
		}
	}()

	return done
}

func startStreaming(ctx context.Context) (err error) {
	cmdArgs := []string{
		"ffmpeg",
		"-i",
		os.Getenv("STREAM"),
		"-v",
		"0",
		"-f",
		"h264",
		"-c",
		"copy",
		"-",
	}

	log.Printf("start streamer process: %s", strings.Join(cmdArgs, " "))

	cmd = exec.Command(cmdArgs[0], cmdArgs[1:]...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return
	}

	if err = cmd.Start(); err != nil {
		return
	}

	log.Printf("streamer process was started: PID %d", cmd.Process.Pid)

	// wait a little till ffmpeg starts write to the stdout
	time.Sleep(100 * time.Millisecond)

	if err = checkProcessState(); err != nil {
		return
	}

	chunk := make([]byte, 0)
	var badChunkCounter int
	chunkLock := sync.Mutex{}

	var minuteTicker *time.Ticker

	save := func() (err error) {
		chunkLock.Lock()

		if len(chunk) == 0 {
			badChunkCounter++

			if badChunkCounter == 5 {
				log.Fatal("last 5 chunks is empty")
			}

			return
		}

		badChunkCounter = 0

		err = saveChunk(chunk)
		if err != nil {
			log.Printf("save chunk failed: %s", err)
		}

		chunk = chunk[:0]

		chunkLock.Unlock()

		return
	}

	go func() {
		minuteTicker = getMinuteTicker()

		err := save()
		if err != nil {
			log.Fatal(err)
		}

		for {
			select {
			case <-ctx.Done():
				return
			case <-minuteTicker.C:
				err = save()
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}()

	packet := make([]byte, 1024*1024)

	for {
		if err = checkProcessState(); err != nil {
			return
		}

		var n int
		n, err = stdout.Read(packet)
		if n == 0 {
			continue
		}
		if err != nil {
			return
		}

		select {
		case <-ctx.Done():
			cmd.Process.Signal(os.Interrupt)
			err = ctx.Err()
			return

		default:
		}

		chunkLock.Lock()
		chunk = append(chunk, packet[:n]...)
		chunkLock.Unlock()

		lastDataTimestamp = time.Now()
	}
}

func checkProcessState() error {
	if lastDataTimestamp.Before(time.Now().Add(-30 * time.Second)) {
		return fmt.Errorf("it seems the process is not alive")
	}

	return nil
}

func getMinuteTicker() *time.Ticker {
	_, _, s := time.Now().Clock()
	secondsSleep := 60 - s
	log.Printf("sleep %d seconds to start chunk ticker", secondsSleep)
	time.Sleep(time.Duration(secondsSleep) * time.Second)
	return time.NewTicker(time.Minute)
}
