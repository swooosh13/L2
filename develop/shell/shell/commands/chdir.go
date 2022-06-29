package commands

import "os"

type ChDir struct{}

func (c *ChDir) Exec(args ...string) (string, error) {
	dir := args[0]
	err := os.Chdir(dir)
	if err != nil {
		return "", err
	}

	dir, err = os.Getwd()
	if err != nil {
		return "", err
	}

	return dir, nil
}


