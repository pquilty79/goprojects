package main

import "fmt"

type triangle struct{
	height float64
	base float64
}
type square struct{
	sideLength float64
}
type shape interface {
	getArea() float64
}

func (t triangle) getArea() float64 {
	return .5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}


func main() {
    t := triangle{height: 4.0, base: 6.5}
	s := square{sideLength: 5}
	printArea(t)
	printArea(s)
}