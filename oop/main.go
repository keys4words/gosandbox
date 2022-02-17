package main

import (
	"fmt"
	"sort"
)

type Employee struct {
	name     string
	position string
	salary   int
	currency string
}

func (e Employee) DisplayInfo() {
	fmt.Println("name: ", e.name)
	fmt.Println("position: ", e.position)
	fmt.Printf("salary: %v %s\n", e.salary, e.currency)
}

type agentSlice []Employee

func (as agentSlice) Len() int           { return len(as) }
func (as agentSlice) Less(i, j int) bool { return as[i].salary < as[j].salary }
func (as agentSlice) Swap(i, j int)      { as[i], as[j] = as[j], as[i] }

func main() {
	emp := Employee{
		name:     "James Bond",
		position: "Agent 007",
		salary:   5500,
		currency: "pounds",
	}
	emp.DisplayInfo()
	emp2 := Employee{
		name:     "Jason Bourn",
		position: "Agent ANB",
		salary:   6500,
		currency: "dollar",
	}
	emp3 := Employee{
		name:     "Bashirov & Petrov",
		position: "Agent KGB",
		salary:   10,
		currency: "wood",
	}
	agents := agentSlice{emp, emp2, emp3}
	sort.Sort(agents)
	fmt.Println(agents)
}
