package taskUtil

import "github.com/cjtoolkit/taskforce"

func CopyFolder(tf *taskforce.TaskForce) func(dest, src string) {
	return func(dest, src string) {
		tf.ExecCmd("cp", "-vRf", src, dest)
	}
}

func CopyFile(tf *taskforce.TaskForce) func(dest string, src ...string) {
	return func(dest string, src ...string) {
		tf.ExecCmd("cp", append(append([]string{"-v"}, src...), dest)...)
	}
}
