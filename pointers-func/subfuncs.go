package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a -b
}

func mult(a, b int) int {
	return a*b
}

func calcAndReturnValidFunc(command string) func(a,b int) int {
	if command == "addition" {
		return func(a,b int) int {return a+b}
	} else if command == "substraction" {
		return func(a,b int) int {return a-b}
	} else {
		return func(a,b int) int {return a*b}
	}
}

func receiveFuncAndReturnValue(f func(a,b int) int) int {
	var intVarA, intVarB int
	intVarA = 100
	intVarB = 200
	return f(intVarA, intVarB)
}

func main() {
	var command string
	command = "substraction"
	res := calcAndReturnValidFunc(command)
	fmt.Println("Result with command: ", command, "value:", res(10,20))
	fmt.Println("func as param: ", receiveFuncAndReturnValue(sub))
}