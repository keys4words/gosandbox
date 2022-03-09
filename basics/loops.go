package main

import "fmt"

func main() {
	loopVar := 0
	for loopVar < 3 {
		fmt.Println("in loop var")
		loopVar++
	}

	for {
		fmt.Println("")
	}
}
