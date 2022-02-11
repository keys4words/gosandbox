package main

import (
	"fmt"
	"packages/rectangle"
)

func init() {
	fmt.Println("Init function for main package")
}

func main() {
	// resAdd := Add(10, 20)
	// resSub := Sub(90, 40)
	// resMult := Mult(50, 4)
	// resDiv := Div(40, 4)
	// fmt.Println(resAdd, resSub, resMult, resDiv)
	r := rectangle.New(10, 20, "green")
	fmt.Println("Perimeter: ", r.Perimeter())

	newR := rectangle.Rectangle{
		A: 10,
		B: 7,
		// color: "red",
	}
	fmt.Println(newR)
}