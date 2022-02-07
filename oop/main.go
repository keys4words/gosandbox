package main

import "fmt"

type Employee struct {
	name string
	position string
	salary int
	currency string
}

func (e Employee) DisplayInfo() {
	fmt.Println("name: ", e.name)
	fmt.Println("position: ", e.position)
	fmt.Printf("salary: %v %s\n", e.salary, e.currency)
}

func main() {
	emp := Employee{
		name: "James Position",
		position:"Agent 007",
		salary: 5500,
		currency: "pounds",
	}
	emp.DisplayInfo()
}