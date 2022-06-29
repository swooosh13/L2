package main

import "fmt"

// интерфейс получателя
type device interface {
	on()
	off()
}

// конкретный получатель
type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
	fmt.Println("TV on")
}

func (t *tv) off() {
	t.isRunning = false
	fmt.Println("TV off")
}