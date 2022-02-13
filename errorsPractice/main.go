package main

import (
	"fmt"
	"strconv"
)

func main() {
	numLiteral := "smth"
	num, err := strconv.Atoi(numLiteral)
	if err != nil {
		fmt.Println("Can not convert this value to int", err)
		return
	}
	fmt.Println("Convertion done:", num)
}
