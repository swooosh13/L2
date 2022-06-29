package commands

import "os"

type Pwd struct {}

func (p *Pwd) Exec(args ...string) (string, error) {
	path, err := os.Getwd()
	return path, err
}
