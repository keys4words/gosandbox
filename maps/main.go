package main

import (
	"fmt"
)


func main() {
	mapper := make(map[string]int)
	mapper["Alice"] = 435
	mapper["Joe"] = 555
	fmt.Println("map: ", mapper)
}