package commands

import (
	"fmt"
	"os"
	"strings"
)

type Shell struct{}

func (s *Shell) RunCommand(command Command, args ...string) string {
	res, err := command.Exec(args...)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		return ""
	}

	return res
}

func (s *Shell) RunCommands(commands []string) {
	pipeBuf := ""
	for _, commandLine := range commands {
		commandLine := strings.Trim(commandLine, " ")
		args := strings.Split(commandLine, " ")
		commandMethod := args[0]

		if len(args) > 1 {
			args = args[1:]
		}

		if commandMethod == "exit" {
			if _, err := fmt.Fprintln(os.Stdout, "exit..."); err != nil {
				fmt.Fprintf(os.Stdout, "error exit comand: %s", err.Error())
				os.Exit(1)
			}
			os.Exit(0)
		}

		command, err := NewCommand(commandMethod)
		if err != nil {
			fmt.Fprintf(os.Stdout, err.Error())
			continue
		}

		args = append(args, pipeBuf)
		pipeBuf = "\n" + s.RunCommand(command, args...)
	}

	fmt.Fprintln(os.Stdout, pipeBuf)
}
