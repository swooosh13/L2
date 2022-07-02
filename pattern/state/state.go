package main

// Состояние описывает общий интерфейс
// для всех конкретных состояний.
type State interface {
	AddItem(int) error
	RequestItem() error
	InsertMoney(money int) error
	DispenseItem() error
}
