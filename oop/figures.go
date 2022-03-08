package main

import (
	"fmt"
	"math"
)

type Figure interface {
	Perimeter() float64
	Info() string
}

type Circle struct {
	Radius int
}

type Rectangle struct {
	Length, Width int
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * float64(c.Radius)
}

func (c Circle) Info() string {
	return fmt.Sprintf("circle with radius %d", c.Radius)
}

func (r Rectangle) Perimeter() float64 {
	return 2 * float64(r.Length+r.Width)
}

func (r Rectangle) Info() string {
	return fmt.Sprintf("rectangle with length:%d and width:%d", r.Length, r.Width)
}

func NewCircle(radius int) *Circle {
	return &Circle{Radius: radius}
}

func main() {
	c := Circle{10}
	// fmt.Println(c.Info())
	// fmt.Printf("Perimeter of c: %.2f\n", c.Perimeter())
	// fmt.Println(Perimeter(c))
	r := Rectangle{10, 20}
	// fmt.Println(r.Info())
	// fmt.Printf("Perimeter of r: %.2f\n", r.Perimeter())
	figures := []Figure{
		c, r,
		NewCircle(30),
		Rectangle{3, 3},
	}
	for _, el := range figures {
		fmt.Println(el.Info())
	}
}
