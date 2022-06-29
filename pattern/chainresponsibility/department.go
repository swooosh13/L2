package main


// интерфейс обработчика
// Handler
type department interface {
	execute(*patient)
	setNext(department)
}
