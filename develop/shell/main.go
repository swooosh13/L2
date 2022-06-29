package main

import (
	"bufio"
	"fmt"
	"os"
	"shell/shell/commands"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	shell := commands.Shell{}

	for {
		fmt.Print("$ ")
		cmdString, err := reader.ReadString('\n')
		cmdString = strings.TrimSuffix(cmdString, "\r\n")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		shell.RunCommands(strings.Split(cmdString, "|"))
	}
}
