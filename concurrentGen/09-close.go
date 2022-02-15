package main

import "fmt"

func generator(ch chan int) {
	for i := 0; i < 25; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go generator(ch)
	// for {
	// 	val, ok := <-ch
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println("Received from channel", val)
	// }
	for val := range ch {
		fmt.Println("Received from channel", val)
	}
}
