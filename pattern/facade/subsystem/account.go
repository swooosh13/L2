package subsystem

import (
	"errors"
	"fmt"
)

type account struct {
	name string
}

func NewAccount(name string) *account {
	return &account{name}
}

func (a *account) checkAccount(name string) error {
	if a.name != name {
		return errors.New("account is incorrect")
	}

	fmt.Println("Account verified successfully")
	return nil
}