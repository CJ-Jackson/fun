package taskUtil

import (
	"bytes"
	"encoding/json"
	"log"
	"os/exec"
	"path"
)

type Env struct {
	GoExe string `json:"GOEXE"`
	GoMod string `json:"GOMOD"`
}

func (e Env) DirPath() string { return path.Dir(e.GoMod) }

func GetEnv() Env {
	buf := &bytes.Buffer{}

	cmd := exec.Command(goCommand(), "env", "-json")
	cmd.Stdout = buf
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Could not run go command: %s", err)
	}

	env := Env{}
	err = json.NewDecoder(buf).Decode(&env)
	if err != nil {
		log.Fatalf("Could not parse json: %s", err)
	}

	if env.GoMod == "" {
		log.Fatal("Not a go module.")
	}

	return env
}
