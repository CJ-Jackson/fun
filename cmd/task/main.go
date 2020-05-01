package main

import (
	"os"

	"github.com/CJ-Jackson/fun/taskUtil"
	"github.com/cjtoolkit/taskforce"
)

func task() *taskforce.TaskForce {
	var (
		tf            = taskforce.InitTaskForce()
		chdir         = taskUtil.Chdir(tf)
		recoverSilent = func() { recover() }
		yarnRun       = taskUtil.YarnRun(tf)
		path          = func() string {
			path := taskUtil.GetEnv().DirPath()
			tf.CheckError(os.Chdir(path + "/asset"))
			return path
		}()
		devEnvironment = true
	)

	tf.Register("yarn", func() {
		tf.ExecCmd("gnode", "yarn")
		tf.ExecCmd("./createLink")
	})

	tf.Register("clean", func() {
		func() {
			defer recoverSilent()
			tf.ExecCmd("rm", "-rf", "live")
		}()

		os.Mkdir("live", 0755)
	})

	{
		// Sass
		const (
			dest = "live/stylesheets/styles.css"
			src  = "dev/sass/styles.scss"
		)

		tf.Register("sass", func() {
			args := []string{"--source-map", "true", "--source-map-contents", "true", "--precision", "8", "--output-style", "compressed"}
			args = append(args, src, dest)
			yarnRun("node-sass", args...)
		})
	}

	tf.Register("rollup", func() {
		env := "BUILD:production"
		if devEnvironment {
			env = "BUILD:development"
		}
		yarnRun("rollup", "-c", "-m", "--environment", env)
	})

	tf.Register("copy", func() {
		copyFolder := taskUtil.CopyFolder(tf)

		copyFolder("live/images", "lfs/images")
		copyFolder("live/fonts", "link/fontawesome/webfonts")
	})

	tf.Register("zip", func() {
		defer chdir(path + "/asset/live")()
		_ = os.Remove("../../asset.zip")

		tf.ExecCmd("zip", "-r0", "../../asset.zip", ".")
	})

	tf.Register("dev", func() {
		tf.Run("yarn", "clean", "sass", "rollup", "copy")
	})

	tf.Register("prod", func() {
		devEnvironment = false
		tf.Run("yarn", "clean", "sass", "rollup", "copy", "zip")
	})

	tf.Register("quick:sass", func() {
		func() {
			defer recoverSilent()
			tf.ExecCmd("rm", "-rf", "live/stylesheets")
		}()
		tf.Run("sass")
	})

	tf.Register("quick:js", func() {
		func() {
			defer recoverSilent()
			tf.ExecCmd("rm", "-rf", "live/javascripts")
		}()
		tf.Run("rollup")
	})

	return tf
}

func main() {
	task().Run(os.Args[1:]...)
}
