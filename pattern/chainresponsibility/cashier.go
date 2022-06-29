package main

import "fmt"

// ConcreteHandler
type cashier struct {
	next department
}

func (c *cashier) execute(p *patient) {
	if p.paymentDone {
		fmt.Println("Payment done")
	}

	// end of chain
	fmt.Println("Cashier getting money from patient")
}

func (c *cashier) setNext(next department) {
	c.next = next
}
