package main

import (
	"fmt"
	"time"
)

func heavyProcess(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

func main() {
	ch := make(chan string)
	go heavyProcess(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received value:", v)
			return
		default:
			fmt.Println("no value received")
		}
	}
}
