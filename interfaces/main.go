package main

import "fmt"

type FigureBuilder interface {
	Area() int
	Perimeter() int
}

type Rectangle struct {
	length, width int
}

func (r Rectangle) Area() int {
	return r.length*r.width
}

func (r Rectangle) Perimeter() int {
	return 2*(r.length+r.width)
}

type Circle struct {
	radius int
}

func (c Circle) Area() int {
	return 3*c.radius*c.radius
}

func (c Circle) Perimeter() int {
	return 2*3*c.radius
}

func main() {
	r1 := Rectangle{10,20}
	r2 := Rectangle{30, 50}
	r3 := Rectangle{1,1}
	c1 := Circle{5}
	c2 := Circle{10}
	c3 := Circle{15}
	// rects := []Rectangle{r1, r2, r3}
	// circles := []Circle{c1, c2, c3}
	// figures := make([]FigureBuilder, 0, 10)
	figures := []FigureBuilder{r1, r2, r3, c1, c2, c3}
	total := 0
	// for _, rect := range rects {
	// 	total += rect.Area()
	// }
	for _, item := range figures {
		total += item.Area()
	}
	fmt.Println("Total area: ", total)
}