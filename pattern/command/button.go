package main

// отправитель
type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}
