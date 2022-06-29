package commands

import "fmt"

type Command interface {
	Exec(args ...string) (string, error)
}

// fabric method
func NewCommand(command string) (Command, error) {
	switch command {
	case "pwd":
		return &Pwd{}, nil
	case "echo":
		return &Echo{}, nil
	case "cd":
		return &ChDir{}, nil
	case "ps":
		return &Ps{}, nil
	case "kill":
		return &Kill{}, nil
	case "exec":
		return &Exec{}, nil
	}

	return nil, fmt.Errorf("command not found: %s", command)
}