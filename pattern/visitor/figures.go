package main

type square struct {
	side int
}

func (s *square) accept(v Visitor) {
	v.visitForSquare(s)
}

func (s *square) getType() string {
	return "square"
}

type circle struct{
	radius int
}

func (c *circle) accept(v Visitor) {
	v.visitForCircle(c)
}

func (c *circle) getType() string {
	return "circle"
}

type rectangle struct {
	l, b int
}

func (r *rectangle) accept(v Visitor) {
	v.visitForRectangle(r)
}

func (r *rectangle) getType() string {
	return "rectangle"
}
