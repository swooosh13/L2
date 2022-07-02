package main

import "fmt"

// Конкретное состояние, когда у нас нет товара в наличии
// работает 1 функция

type NoItemState struct {
	vendingMachine *vendingMachine
}

func (i *NoItemState) RequestItem() error {
	return fmt.Errorf("Item out of stock")
}

func (i *NoItemState) AddItem(count int) error {
	i.vendingMachine.incrementItemCount(count)
	// задаем что товар есть в наличии - меняем state
	i.vendingMachine.SetState(i.vendingMachine.hasItem)
	return nil
}

func (i *NoItemState) InsertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}
func (i *NoItemState) DispenseItem() error {
	return fmt.Errorf("Item out of stock")
}
