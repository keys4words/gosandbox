package main

import (
	"fmt"
)

func ReadChan(ch chan int) {
	value := <-ch
	fmt.Println("Chan value:", value)
}

func main() {
	fmt.Println("Start main...")
	var ch chan int
	ch = make(chan int, 1)
	ch <- 100
	go ReadChan(ch)
	// time.Sleep(1 * time.Second)
	ch <- 34
	fmt.Println("End main")
}
