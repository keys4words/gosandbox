package main

import "fmt"

type Worker interface {
	Work()
}

type Employee struct {
	name string
	age int
}

func (e Employee) Work() {
	fmt.Println("Now Employee with name", e.name, "is working")
}

func Describer(w Worker) {
	fmt.Printf("Interface with type %T and value %v\n", w, w)
}

type Student struct {
	name string
	courseNumber int
}

func (s Student) WantToEat() {
	fmt.Println("Student with name", s.name, "want to eat!")
}

func (s Student) Work() {
	fmt.Println("Student with name", s.name, "is working")
}

func main() {
	emp := Employee{"James", 44}
	var workerEmployee Worker = emp
	// workerEmployee.Work()
	// Describer(workerEmployee)

	std := Student{"Johny", 4444}
	var workerStudent Worker = std
	// workerStudent.Work()
	// Describer(workerStudent)

	var workers []Worker = []Worker{workerStudent, workerEmployee}
	for _, worker := range workers {
		Describer(worker)
	}
}