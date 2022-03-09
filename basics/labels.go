package main

import "fmt"

func main() {
outer:
	for i := 0; i < 5; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i:%d and j:%d give sum i+j=%d\n", i, j, i+j)
			if i == j {
				break outer
			}
		}
	}
}
