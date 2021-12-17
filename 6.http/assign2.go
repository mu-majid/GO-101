package main

import "fmt"

type shape interface {
	// getArea() interface{} // empty interface to allow any value (here it could be float or int)
	getArea() float64 // empty interface to allow any value (here it could be float or int)

}

type square struct {
	sideLength float64
}

type triangle struct {
	base   float64
	height float64
}

func mainFunc() {
	sq := square{sideLength: 5}
	tr := triangle{base: 5, height: 6}

	printArea(sq)
	printArea(tr)

}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
