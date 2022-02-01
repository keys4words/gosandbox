package main

import (
	"fmt"
	// "bufio"
	// "os"
	"strconv"
)

func main() {
	// problems when you need several words input
	// var a string
	// // fmt.Scan(&a)
	// // fmt.Println(a)
	// input := bufio.NewScanner(os.Stdin)
	// // input.Scan()	//catch till end of the string
	// if input.Scan() {
	// 	a = input.Text()
	// }
	// fmt.Println(a)

	// fmt.Println("For loop started")
	// for {
	// 	if input.Scan() {
	// 		result := input.Text()
	// 		if result == "" {
	// 			break
	// 		}
	// 		fmt.Println("input string is: ", result)
	// 	}
	// }

	// convert string to smth
	numStr := "10"
	numInt, _ := strconv.Atoi(numStr)
	fmt.Printf("%d is %T\n", numInt, numInt)
	numInt64, _ := strconv.ParseInt(numStr, 10, 64)
	fmt.Println(numInt64)
}