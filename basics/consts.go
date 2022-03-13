package main

import (
	"fmt"
	// "math"
)

const (
	IP_ADDRESS = "127.0.0.5"
	PORT = "5432"
	ADMIN_EMAIL string = "admin@mail.com"
)


func main() {
	const a = 10
	// a = 20
	fmt.Println(a)
	fmt.Println(IP_ADDRESS, ADMIN_EMAIL)

	// const sqrt = math.Sqrt(25)
	// fmt.Println("Var sqrt: ", sqrt)
	const NUMERIC = 100
	var numInt8 int8 = NUMERIC
	var numInt32 int32 = NUMERIC
	var numInt64 int64 = NUMERIC
	fmt.Printf("Value %v and Type %T\n", numInt8, numInt8)
	fmt.Printf("Value %v and Type %T\n", numInt32, numInt32)
	fmt.Printf("Value %v and Type %T\n", numInt64, numInt64)
}