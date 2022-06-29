package main


// интерфейс команды
type command interface {
	execute()
}

// конкретная команда
type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

// конкретная комманда
type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

