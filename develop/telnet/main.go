package main

import (
	app2 "telnet/internal/app"
	"telnet/internal/telnet"
)

func main() {
	flags := app2.GetFlags()
	client := telnet.NewClient(flags.Host, flags.Port, flags.Timeout)

	app := app2.NewApp(flags, client)
	app.Run()
}
