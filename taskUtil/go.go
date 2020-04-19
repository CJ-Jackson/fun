package taskUtil

import (
	"os"
	"os/exec"

	"github.com/cjtoolkit/taskforce"
)

func goCommand() string {
	goCmd := os.Getenv("GOCMD")
	if goCmd != "" {
		return goCmd
	}

	return "go"
}

func GoBuild(tf *taskforce.TaskForce) func(output, src string) {
	return func(output, src string) {
		cmd := exec.Command(goCommand(), "build", "-v", "-o", output, src)
		cmd.Env = append(os.Environ(), "GOOS=linux", "GOARCH=amd64")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		tf.CheckError(cmd.Run())
	}
}
