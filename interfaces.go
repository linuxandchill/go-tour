package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	area() float64
}

//Square implements Shape interface
func (s Square) area() float64 {
	return s.side * s.side
}

//Circle implements Shape interface
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

//asking for a type Shape which is an iface
func info(z Shape) float64 {
	fmt.Println(z)
	//call area func on each shape
	return (z.area())
}

func main() {
	c := Circle{4}
	s := Square{3}
	circle_area := info(c)
	square_area := info(s)
	fmt.Println(circle_area)
	fmt.Println(square_area)
}
