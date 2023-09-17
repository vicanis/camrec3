package stream

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

var cmd *exec.Cmd

func Start(ctx context.Context) chan error {
	done := make(chan error, 1)

	go func() {
		done <- startStreaming(ctx)
	}()

	return done
}

func startStreaming(ctx context.Context) (err error) {
	if os.Getenv("STREAM") == "" {
		err = errors.New("no stream URL")
		return
	}

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
	chunkLock := sync.Mutex{}

	var minuteTicker *time.Ticker

	saveChunk := func() (err error) {
		chunkLock.Lock()

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

		chunk = make([]byte, 0)

		chunkLock.Unlock()

		return
	}

	go func() {
		minuteTicker = getMinuteTicker()

		err := saveChunk()
		if err != nil {
			log.Fatal(err)
		}

		for {
			select {
			case <-ctx.Done():
				return
			case <-minuteTicker.C:
				err = saveChunk()
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}()

	for {
		if err = checkProcessState(); err != nil {
			return
		}

		packet := make([]byte, 1024*1024)

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
	}
}

func checkProcessState() error {
	state := cmd.ProcessState

	if state == nil {
		return nil
	}

	if state.Exited() {
		return fmt.Errorf("process exited with code %d", state.ExitCode())
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
