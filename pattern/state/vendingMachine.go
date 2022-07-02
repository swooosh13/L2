package main

import "fmt"

// содержит ссылку на один из объектов состояний и делегировать ему работу
//

type vendingMachine struct {
	hasItem       State // состояние
	itemRequested State
	hasMoney      State
	noItem        State

	currentState State // конкретное состояние

	itemCount int
	itemPrice int
}

func NewVendingMachine(itemCount, itemPrice int) *vendingMachine {
	v := &vendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	hasItemState := &HasItemState{
		v,
	}
	itemRequestedState := &ItemRequestedState{
		v,
	}
	hasMoneyState := &HasMoneyState{
		v,
	}
	noItemState := &NoItemState{
		v,
	}

	v.SetState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}

func (v *vendingMachine) RequestItem() error {
	return v.currentState.RequestItem()
}

func (v *vendingMachine) AddItem(count int) error {
	fmt.Printf("Adding %d items\n", count)
	return v.currentState.AddItem(count)
}

func (v *vendingMachine) InsertMoney(money int) error {
	return v.currentState.InsertMoney(money)
}

func (v *vendingMachine) DispenseItem() error {
	return v.currentState.DispenseItem()
}

func (v *vendingMachine) SetState(state State) {
	v.currentState = state
}

func (v *vendingMachine) incrementItemCount(count int) {
	v.itemCount += count
}
