package main

import "fmt"

type HasItemState struct {
	vendingMachine *vendingMachine
}

func (s *HasItemState) RequestItem() error {
	if s.vendingMachine.itemCount == 0 {
		s.vendingMachine.SetState(s.vendingMachine.noItem)
		return fmt.Errorf("No item present")
	}

	fmt.Println("Item requested")
	s.vendingMachine.SetState(s.vendingMachine.itemRequested)
	return nil
}

func (s *HasItemState) AddItem(count int) error {
	fmt.Printf("%d items added\n", count)
	s.vendingMachine.incrementItemCount(count)
	return nil
}

func (s *HasItemState) InsertMoney(money int) error {
	return fmt.Errorf("Please select item first")
}
func (s *HasItemState) DispenseItem() error {
	return fmt.Errorf("Please select item first")
}
