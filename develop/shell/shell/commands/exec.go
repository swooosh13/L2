package commands

import (
	"fmt"
	"os"
	"os/exec"
)

type Exec struct {}

func (e *Exec) Exec(args ...string) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("no arguments for command exepsc")
	}

	cm := exec.Command(args[0], args[1:]...)
	cm.Stdin = os.Stdin
	cm.Stdout = os.Stdout
	cm.Stderr = os.Stderr
	return "", cm.Run()
}