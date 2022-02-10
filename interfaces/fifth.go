package main

import "fmt"

type Describer interface {
	Describe()
}

type Student struct {
	name string
	age int
}

func (std Student) Describe() {
	fmt.Printf("Student name: %s and age: %d\n", std.name, std.age)
}

type Professor struct {
	name string
	age int
}

func (prof *Professor) Describe() {
	fmt.Printf("Professor name: %s and age: %d\n", prof.name, prof.age)
}


func main() {
	var desc Describer
	std1 := Student{"John", 21}
	desc = std1
	desc.Describe()
	std2 := Student{"Bob", 22}
	desc = std2
	desc.Describe()

	var desc2 Describer
	prof1 := &Professor{"Niels Bohr", 89}
	desc2 = prof1
	desc2.Describe()
}