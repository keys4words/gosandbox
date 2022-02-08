package main

import "fmt"

// func (a *int) IsEven() bool {
// 	if a%2 == 0 {
// 		return true
// 	}
// 	return false
// }

type MySuperInt int
func (x *MySuperInt) IsEven() bool {
	if *x%2 == 0 {
		return true
	}
	return false
}

func main() {
	num1 := MySuperInt(100)
	num2 := MySuperInt(101)
	fmt.Println(num1.IsEven())
	fmt.Println(num2.IsEven())
}