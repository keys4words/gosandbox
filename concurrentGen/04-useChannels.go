package main

import "fmt"

func newGoRoutine(done chan bool) {
	fmt.Println("Hey, I'm new GoRoutine!")
	done <- true
}
func main() {
	done := make(chan bool)
	go newGoRoutine(done)
	<-done
	fmt.Println("Main goroutine work!")
}
