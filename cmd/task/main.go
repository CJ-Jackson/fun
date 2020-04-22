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
	)

	tf.Register("yarn", func() {
		tf.ExecCmd("./yarn")
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

		tf.Register("sass:dev", func() {
			args := []string{"--source-map", "true", "--source-map-contents", "true", "--precision", "8"}
			args = append(args, src, dest)
			yarnRun("node-sass", args...)
		})

		tf.Register("sass:prod", func() {
			args := []string{"--source-map", "true", "--source-map-contents", "true", "--precision", "8", "--output-style", "compressed"}
			args = append(args, src, dest)
			yarnRun("node-sass", args...)
		})
	}

	tf.Register("rollup", func() {
		yarnRun("rollup", "-c", "-m")
	})

	tf.Register("copy", func() {
		copyFolder := taskUtil.CopyFolder(tf)

		copyFolder("live/images", "lfs/images")
		copyFolder("live/fonts", "link/fontawesome/webfonts")
	})

	tf.Register("uglify:prod", func() {
		yarnRun("uglifyjs", "live/javascripts/bundle.js",
			"-o", "live/javascripts/bundle.js", "--source-map", "content='live/javascripts/bundle.js.map',url='bundle.js.map'", "--compress")
	})

	tf.Register("zip", func() {
		defer chdir(path + "/asset/live")()
		_ = os.Remove("../../asset.zip")

		tf.ExecCmd("zip", "-r0", "../../asset.zip", ".")
	})

	tf.Register("dev", func() {
		tf.Run("yarn", "clean", "sass:dev", "rollup", "copy")
	})

	tf.Register("prod", func() {
		tf.Run("yarn", "clean", "sass:prod", "rollup", "copy", "zip")
	})

	tf.Register("quick:sass", func() {
		func() {
			defer recoverSilent()
			tf.ExecCmd("rm", "-rf", "live/stylesheets")
		}()
		tf.Run("sass:dev")
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
