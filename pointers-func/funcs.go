package main

import "fmt"

func add(a int, b int) int {
	result := a+b
	return result
}

func mult(a,b,c,d int) int {
	return a*b*c*d
}

func rectangleParams(length, width float64) (float64, float64) {
	var perimeter = 2*(length+width)
	var area = length*width
	return perimeter, area
}

func namedReturn(a,b int) (perimeter, area int) {
	perimeter = 2*(a+b)
	area = a*b
	return
}

func funcWithReturn(a,b int) (int, bool) {
	if a>b {
		return a-b, true
	}
	if a==b {
		return a, true
	}
	return 0, false
}

//show param func
func emptyReturn(a int) {
	fmt.Println("Empty Return with param: ", a)
}

func main() {
	// fmt.Println("Hey, dude!")
	res := add(10, 20)
	fmt.Println(res)
	fmt.Println(mult(1,2,3,4))
	per, area := rectangleParams(10.5, 10.5)
	fmt.Println("Area = ", area)
	fmt.Println("Perimeter = ", per)
	per1, area1 := namedReturn(10,10)
	fmt.Println("NamedReturn, per & area: ", per1, area1)
	emptyReturn(10)
}