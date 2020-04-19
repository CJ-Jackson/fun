package taskUtil

import (
	"fmt"
	"io"
	"os"

	"github.com/cjtoolkit/taskforce"
)

func Concat(tf *taskforce.TaskForce) func(dest string, src ...string) {
	return func(dest string, src ...string) {
		destFile, err := os.Create(dest)
		tf.CheckError(err)
		defer destFile.Close()

		for _, srcFilename := range src {
			fmt.Printf("Concat: '%s' -> '%s'", srcFilename, dest)
			fmt.Println()
			srcFile, err := os.Open(srcFilename)
			tf.CheckError(err)

			_, _ = io.Copy(destFile, srcFile)
			_ = srcFile.Close()
		}
	}
}
