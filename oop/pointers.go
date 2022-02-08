package main

import "fmt"

type Employee struct {
	name string
	salary int
}

//method to copy employee instance
func (e Employee) SetName(newName string){
	e.name = newName
}
//method to use anchor with instance
func (e *Employee) SetSalary(newSalary int){
	e.salary = newSalary
}

func main() {
	e := Employee{"Alex", 3000}
	fmt.Println("Before setting params: ", e)
	e.SetName("Bob")
	// fmt.Println("After setting params: ", e)
	e.SetSalary(6000)
	fmt.Println("Afther setting params: ", e)
}