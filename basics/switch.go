package main

import "fmt"

func main() {
	var price int
	fmt.Scan(&price)
	switch {
	case price < 100:
		fmt.Println("First case")
	case price > 100 && price < 1000:
		fmt.Println("Second case")
	default:
		fmt.Println("Default case")
	}
	// var ageGroup string = "b"
	// switch ageGroup {
	// case "a", "b", "c":
	// 	fmt.Println("age group 10-30")
	// case "d", "e":
	// 	fmt.Println("age 40-50")
	// }

	var num int
	fmt.Scan(&num)
	switch num {
	case 100:
		fmt.Println("-> 100")
		fallthrough
	case 300:
		fmt.Println("-> 300")
	}
}
