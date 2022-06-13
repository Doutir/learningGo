package main

import "fmt"

type square struct {
	sideLength float64
}
type triangule struct {
	height float64
	base   float64
}
type shape interface {
	getArea() float64
}

func main() {
	t := triangule{
		height: 10.03,
		base:   23,
	}
	s := square{
		sideLength: 12,
	}
	printArea(t)
	printArea(s)

}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangule) getArea() float64 {
	return 0.5 * t.base * t.height
}
