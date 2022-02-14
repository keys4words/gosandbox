package main

import (
	"fmt"
	"time"
)

func newGoRoutine(done chan bool) {
	fmt.Println("Hey, i'm newGoRoutine and I'm going sleep!")
	time.Sleep(4 * time.Second)
	fmt.Println("newGoRoutine awake and going to send data to channel")
	done <- true
}
func main() {
	done := make(chan bool)
	fmt.Println("I'm main goroutine and i wanna call newGoRoutine")
	go newGoRoutine(done)
	<-done
	fmt.Println("OK, main goroutine received data and gonna die!")
}
