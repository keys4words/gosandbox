package main

import "fmt"

func sumAll(args ...int) int {
	sums := 0
	for _, el := range args {
		sums += el
	}
	return sums
}

func main() {
	fmt.Println(sumAll(1, 2, 3, 4, 5))
	fmt.Println(sumAll())
}
