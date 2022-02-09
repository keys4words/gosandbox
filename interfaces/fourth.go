package main

import "fmt"

func DoSomething(pretedent interface{}) {
	switch pretedent.(type) {
	case string:
		fmt.Println("This is string")
	case int:
		fmt.Println("This is int")
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	DoSomething(10)
	DoSomething("another string")
	DoSomething(func(a, b int) int {return a+b})
}