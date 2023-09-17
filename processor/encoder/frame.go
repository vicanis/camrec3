package encoder

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func ExtractFrame(source []byte, offset int) (frame []byte, err error) {
	tempFiles := make([]string, 0)

	defer func() {
		for _, f := range tempFiles {
			if err := os.Remove(f); err != nil {
				log.Printf("remove temp file failed: %s", err)
			}
		}
	}()

	tmp, err := os.CreateTemp(os.TempDir(), "src")
	if err != nil {
		return
	}

	tempFiles = append(tempFiles, tmp.Name())

	_, err = tmp.Write(source)
	if err != nil {
		return
	}

	tmp.Close()

	tempFiles = append(tempFiles, tmp.Name()+".jpg")

	cmdArgs := []string{
		"ffmpeg",
		"-i",
		tmp.Name(),
		"-frames:v",
		"1",
		"-ss",
		fmt.Sprint(offset),
		"-t",
		"60",
		"-y",
		tmp.Name() + ".jpg",
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
		log.Printf("process output: %s", stdout.String())
		err = fmt.Errorf("exit code: %d", exitCode)
		return
	}

	return os.ReadFile(tmp.Name() + ".jpg")
}
