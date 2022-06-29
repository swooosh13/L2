package main

import "fmt"

// Laptop - конкретный продукт
type Laptop struct {
	Type    string
	Core    int
	Mem     int
	Monitor bool
}

func NewLaptop() Computer {
	return &Laptop{
		Type:    LaptopType,
		Core:    4,
		Mem:     8,
		Monitor: true,
	}
}

func (s *Laptop) GetType() string {
	return s.Type
}

func (s *Laptop) PrintDetails() {
	fmt.Printf("%s Core:[%d], Mem:[%d], Monitor:[%v]\n", s.Type, s.Core, s.Mem, s.Monitor)
}

