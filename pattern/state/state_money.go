package main

import "fmt"

type HasMoneyState struct {
	vendingMachine *vendingMachine
}

func (s *HasMoneyState) RequestItem() error {
	return fmt.Errorf("Item already requested")
}

func (s *HasMoneyState) AddItem(count int) error {
	return fmt.Errorf("Item Dispense in progress")
}

func (s *HasMoneyState) InsertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}
func (s *HasMoneyState) DispenseItem() error {
	fmt.Println("Dispensing item")
	s.vendingMachine.itemCount--
	if s.vendingMachine.itemCount <= 0 {
		s.vendingMachine.SetState(s.vendingMachine.noItem)
	} else {
		s.vendingMachine.SetState(s.vendingMachine.hasItem)
	}

	return nil
}
