package main

import "fmt"

func bar() {
	panic("bar panic")
}

func foo() {
	defer func() {
		recover()
	}()
	for i := 0; i < 10; i++ {
		fmt.Printf("foo before cycle #%d\n", i)
		// bar()
		go func() {
			bar()
		}()
		fmt.Printf("foo after cycle #%d\n", i)
	}
}

func main() {
	foo()
}
