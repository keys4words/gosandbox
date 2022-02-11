package rectangle

import "fmt"

var (
	name string = "John"	
	age int = 34
)

func init() {
	fmt.Println("Init function from rectangle package")
	fmt.Println("Name:", name, "Age:", age)
}

type Rectangle struct {
	A,B int
	color string
}

func New(newA int, newB int, newColor string) *Rectangle {
	return &Rectangle{newA, newB, newColor}
}

func (r *Rectangle) Perimeter() int {
	return 2*(r.A+r.B)
}