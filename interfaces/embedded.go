package main

import "fmt"

type PerimeterCalc interface {
	Perimeter() int
}

type AreaCalc interface {
	Area() int
}

type Rectangle struct {
	a, b int
	color string
}

func (r Rectangle) Perimeter() int {
	return 2*(r.a+r.b)
}

func (r Rectangle) Area() int {
	return r.a*r.b
}

type FigureFeatureCalc interface {
	PerimeterCalc
	AreaCalc
}


func main() {
	var pCalc PerimeterCalc
	var aCalc AreaCalc
	rect := Rectangle{10, 20, "green"}
	pCalc = rect
	fmt.Println("Perimeter:", pCalc.Perimeter())
	aCalc = rect
	fmt.Println("Area:", aCalc.Area())

	var ffCalc FigureFeatureCalc
	ffCalc = rect
	fmt.Println(ffCalc.Area())
	fmt.Println(ffCalc.Perimeter())
}