package rectangle

import "fmt"

func init() {
	fmt.Println("Init function from rectangle package")
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