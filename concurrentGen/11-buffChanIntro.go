package main

import "fmt"

func main() {
	ch := make(chan string, 2)
	ch <- "James"
	ch <- "Bond"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
