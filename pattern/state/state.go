package main

// Состояние описывает общий интерфейс
// для всех конкретных состояний.
type state interface {
	AddItem(int) error
	RequestItem() error
}
