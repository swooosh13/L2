package app

import (
	"bufio"
	"log"
	"os"
	"os/signal"
	"syscall"
	"telnet/internal/telnet"
)

type App struct {
	flags  *Flags
	client *telnet.Client
}

func NewApp(flags *Flags, client *telnet.Client) *App {
	return &App{
		flags:  flags,
		client: client,
	}
}

func (a *App) Run() {
	err := a.client.Connect()
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer func() {
		if err := a.client.Close(); err != nil {
			log.Fatalf(err.Error())
		}
	}()

	errors := make(chan error, 1)
	go a.Receiver(errors)
	go a.Sender(errors)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-done:
			return
		case err := <-errors:
			log.Printf("error: %s", err.Error())
			return
		}
	}
}

func (a *App) Receiver(errors chan error) {
	for {
		if err := a.client.Receive(); err != nil {
			errors <- err
		}
	}
}

func (a *App) Sender(errors chan error) {
	reader := bufio.NewReader(os.Stdin)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			errors <- err
		}

		if err := a.client.Send(msg); err != nil {
			errors <- err
		}
	}
}
