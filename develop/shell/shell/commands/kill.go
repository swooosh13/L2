package commands

import (
	"fmt"
	"os"
	"strconv"
)

type Kill struct {}

func (k *Kill) Exec(args ...string) (string, error) {
	pid, err := strconv.Atoi(args[0])
	if err != nil {
		return "", fmt.Errorf("error converting: %s", err.Error())
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		return "", fmt.Errorf("erocess not found: %s", err.Error())
	}

	if err = process.Kill(); err != nil {
		return "", fmt.Errorf("error killing process: %s", err.Error())
	}

	return fmt.Sprintf("PID:%d killed successfully", pid), nil
}
