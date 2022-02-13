package main

import (
	"fmt"
	"time"
)

func newGoRoutine() {
	fmt.Println("Hey, I'm new GoRoutine!")
}

func main() {
	go newGoRoutine()
	time.Sleep(1 * time.Second)
	fmt.Println("Main goroutine work")
}
