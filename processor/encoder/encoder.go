package encoder

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Encode(source []byte, offset int) (chunk []byte, err error) {
	tempFiles := make([]string, 0)

	defer func() {
		for _, f := range tempFiles {
			if err := os.Remove(f); err != nil {
				log.Printf("remove temp file failed: %s", err)
			}
		}
	}()

	src, err := os.CreateTemp(os.TempDir(), "src")
	if err != nil {
		return
	}

	tempFiles = append(tempFiles, src.Name())

	_, err = src.Write(source)
	if err != nil {
		return
	}

	src.Close()

	dst, err := os.CreateTemp(os.TempDir(), "dst")
	if err != nil {
		return
	}

	tempFiles = append(tempFiles, dst.Name())

	dst.Close()

	cmdArgs := []string{
		"ffmpeg",
		"-i",
		src.Name(),
		"-f",
		"mp4",
		"-vcodec",
		"copy",
		"-v",
		"0",
		"-ss",
		fmt.Sprint(offset),
		"-t",
		"30",
		"-y",
		dst.Name(),
	}

	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)

	stdout := &strings.Builder{}

	cmd.Stdout = stdout
	cmd.Stderr = stdout

	err = cmd.Start()
	if err != nil {
		return
	}

	cmd.Wait()

	exitCode := cmd.ProcessState.ExitCode()

	if exitCode != 0 {
		err = fmt.Errorf("exit code: %d", exitCode)
		log.Print(err)
		return
	}

	chunk, err = os.ReadFile(dst.Name())

	return
}
