package main

import "fmt"

func main() {
	var a chan int
	if a == nil {
		fmt.Println("channel is nil, let's define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T\n", a)
	}

}