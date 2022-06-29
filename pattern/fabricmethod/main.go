package main

import "fmt"

var types = []string{ServerType, LaptopType, PersonalComputerType}

func main() {
	for _, t := range types {
		computer := NewComputer(t)
		computer.PrintDetails()
		fmt.Println(computer.GetType())
	}
}

