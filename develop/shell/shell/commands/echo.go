package commands

import (
	"os/exec"
)

type Echo struct {}

func (e *Echo) Exec(args ...string) (string, error) {
	data, err := exec.Command("echo", args...).Output()
	return string(data), err
}
