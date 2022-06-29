package main

import (
	"fmt"
)

type middleCoordinates struct {
	x int
	y int
}

func (a *middleCoordinates) visitForSquare(s *square) {
	a.x = s.side / 2
	a.y = s.side / 2
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *middleCoordinates) visitForCircle(c *circle) {
	a.x = c.radius
	a.y = c.radius
	fmt.Println("Calculating middle point coordinates for circle")
}
func (a *middleCoordinates) visitForRectangle(t *rectangle) {
	a.x = t.l / 2
	a.y = t.b / 2
	fmt.Println("Calculating middle point coordinates for rectangle")
}
