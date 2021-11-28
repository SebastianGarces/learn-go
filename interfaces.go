package main

import (
	"fmt"
	"math"
)

type Square struct {
	length float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	area() float64
	circumf() float64
}

func (s Square) area() float64 {
	return s.length * s.length
}

func (s Square) circumf() float64 {
	return s.length * 4
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeIngo(shape Shape) {
	fmt.Printf("%T area: %0.2f \n", shape, shape.area())
	fmt.Printf("%T circumf: %0.2f \n", shape, shape.circumf())
}

func main() {
	shapes := []Shape{
		Circle{radius: 5},
		Circle{radius: 10},
		Square{length: 4.5},
		Square{length: 2},
		Square{length: 22},
	}

	for _, v := range shapes {
		printShapeIngo(v)
		fmt.Println("-------------------")
	}
}
