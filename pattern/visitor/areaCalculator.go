package main

import (
	"fmt"
	"math"
)


// конкретный посетитель

type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *square) {
	a.area += s.side * s.side
	fmt.Println("Calculating area for square")
}
func (a *areaCalculator) visitForCircle(c *circle) {
	a.area += int(float64(c.radius*c.radius) * math.Pi)
	fmt.Println("Calculating area for circle")
}
func (a *areaCalculator) visitForRectangle(r *rectangle) {
	a.area += r.b * r.l
	fmt.Println("Calculating area for rectangle")
}


