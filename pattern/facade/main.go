package main

import (
	"facade/subsystem"
	"fmt"
	"log"
)

func main() {
	accountId := "23flk3"
	code := 1722

	fmt.Println()
	ws := subsystem.NewWalletFacade(accountId, code)
	fmt.Println()

	if err := ws.AddMoneyToWallet(accountId, code, 16); err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println()

	if err := ws.DeductMoneyFromWallet(accountId, code, 10); err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println()
}
