package main

import "fmt"

type Empty interface{}

func Describer(pretedent interface{}){
	fmt.Printf("Type = %T and value %v\n", pretedent, pretedent)
}

type Student struct {
	name string
}

func SemiGeneric(pretedent interface{}) {
	val, ok := pretedent.(int)
	fmt.Println("Value:", val, "OK?:", ok)
}

func main() {
	strSample := "Hello World!"
	Describer(strSample)
	intSample := 300
	Describer(intSample)
	Describer(Student{"Peter Penn"})
	SemiGeneric(10)
	SemiGeneric(Student{"John Doe"})
	SemiGeneric("some string")
}