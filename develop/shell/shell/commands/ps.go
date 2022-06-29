package commands

import "os/exec"

type Ps struct {}

func (p *Ps) Exec(args ...string) (string, error) {
	res, err := exec.Command("ps").Output()
	return string(res), err
}
