package taskUtil

import "github.com/cjtoolkit/taskforce"

func YarnRun(tf *taskforce.TaskForce) func(name string, args ...string) {
	return func(name string, args ...string) {
		args = append([]string{"run", name}, args...)
		tf.ExecCmd("./yarn", args...)
	}
}
