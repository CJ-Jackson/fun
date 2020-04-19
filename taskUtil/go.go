package taskUtil

import (
	"os"
)

func goCommand() string {
	goCmd := os.Getenv("GOCMD")
	if goCmd != "" {
		return goCmd
	}

	return "go"
}
