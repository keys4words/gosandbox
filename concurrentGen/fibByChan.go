package main

import (
	"fmt"
	"time"
)

func fib(n int, c chan int) {
	a, b := 1, 1
	for i := 0; i < n; i++ {
		c <- a
		a, b = b, a+b
		time.Sleep(1 * time.Second)
	}
	close(c)
}

func main() {
	c := make(chan int)
	go fib(10, c)
	for i := range c {
		fmt.Println(i)
	}
}