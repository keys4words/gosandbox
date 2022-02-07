package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Rectangle struct {
	length, width int
}

func Perimeter(c Circle) float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}


func main() {
	c := Circle{10.5}
	fmt.Println(c.Perimeter())
	fmt.Println(Perimeter(c))
	r := Rectangle{10, 20}
	fmt.Println("Rectangle: ", r.Perimeter())
}