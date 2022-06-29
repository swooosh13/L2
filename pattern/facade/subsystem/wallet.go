package subsystem

import (
	"errors"
	"fmt"
)

type wallet struct {
	balance int
}

func NewWallet() *wallet {
	return &wallet{
		balance: 0,
	}
}

func (w *wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Println("Wallet balance added successfully")
}

func (w *wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return errors.New("balance is not sufficient")
	}

	fmt.Println("wallet balance is sufficient")
	w.balance = w.balance - amount
	return nil
}