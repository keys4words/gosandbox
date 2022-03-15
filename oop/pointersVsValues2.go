package main

import "fmt"

type Rectangle struct {
	length, width int
}

func (r *Rectangle) Area() int {
	return r.length*r.width
}

func Area(r *Rectangle) int {
	return r.length*r.width
}

func (r *Rectangle) SetWidth(newWidth int) {
	r.width = newWidth
}


func main() {
	rectAsValue := Rectangle{10,10}
	rectAsPointer := &rectAsValue
	fmt.Println("Call Area func for &rectAsPointer: ", Area(rectAsPointer))
	fmt.Println("Call Area method for &rectAsPointer: ", rectAsPointer.Area())
	fmt.Println("Call Area method for rectAsValue: ", rectAsValue.Area())
	// fmt.Println("Call Area func for rectAsValue: ", Area(rectAsValue))
	fmt.Println("Before changing width: ", rectAsValue)
	rectAsPointer.SetWidth(999)
	fmt.Println("After change via method SetWidth for &rect: ", rectAsValue)
	rectAsValue.SetWidth(888)
	fmt.Println("After change via method SetWidth for rect: ", rectAsValue)

}