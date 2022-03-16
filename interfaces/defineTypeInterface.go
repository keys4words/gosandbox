package main

import "fmt"

type Describer interface {
	Describe()
}

type Student struct {
	name string
	age  int
}

func (std Student) Describe() {
	fmt.Printf("Student %s and %d\n", std.name, std.age)
}

func DoSomething(pretedent interface{}) {
	switch v := pretedent.(type) {
	case string:
		fmt.Println("This is string")
	case int:
		fmt.Println("This is int")
	case Describer:
		fmt.Println("This is describer")
		v.Describe()
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	DoSomething(10)
	DoSomething("another string")
	DoSomething(func(a, b int) int { return a + b })
	std := Student{"Alex", 22}
	DoSomething(std)
}
