package main

import "fmt"

// PersonalComputer - конкретный продукт
type PersonalComputer struct {
	Type    string
	Core    int
	Mem     int
	Monitor bool
}

func NewPersonalComputer() Computer {
	return &PersonalComputer{
		Type:    PersonalComputerType,
		Core:    8,
		Mem:     16,
		Monitor: true,
	}
}

func (s *PersonalComputer) GetType() string {
	return s.Type
}

func (s *PersonalComputer) PrintDetails() {
	fmt.Printf("%s Core:[%d], Mem:[%d], Monitor:[%v]\n", s.Type, s.Core, s.Mem, s.Monitor)
}
