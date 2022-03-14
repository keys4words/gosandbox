package main

import "fmt"

//1. Структура - заименованный набор полей (свойств - состояний), опредляющий новый тип данных
type Student struct {
	firstName string
	lastName  string
	age       int
}

type AnotherStudent struct {
	firstName, lastName, groupName string
	age, courceNumber              int
}

func PrintStudent(std Student) {
	fmt.Println("==============")
	fmt.Println("First Name: ", std.firstName)
	fmt.Println("Last Name: ", std.lastName)
	fmt.Println("Age: ", std.age)
}

//stuct with anonymous properties
type Human struct {
	firstName string
	string
	bool
}

func main() {
	stud1 := Student{
		firstName: "Ivan",
		lastName:  "Petrov",
		age:       21,
	}
	fmt.Println(stud1)
	PrintStudent(stud1)
	stud2 := Student{
		firstName: "John",
	}
	PrintStudent(stud2)

	//Anonymous struct
	anonStudent := struct {
		age           int
		groupID       int
		professorName string
	}{
		age:           23,
		groupID:       1,
		professorName: "Alexeev",
	}
	fmt.Println("Anonymous struct: ", anonStudent)

	//access to properties
	studNova := Student{"Vova", "Petrov", 19}
	fmt.Println(studNova.firstName)
	studNova.age += 10
	fmt.Println(studNova.age)

	//initialization of empty struct
	emptyStud1 := Student{}
	var emptyStud2 Student
	PrintStudent(emptyStud1)
	PrintStudent(emptyStud2)

	// pointer to struct instance
	studPointer := &Student{
		firstName: "Mike",
		lastName:  "Tyson",
		age:       47,
	}
	fmt.Println("Value studPointer: ", studPointer)
	secondPointer := studPointer
	(*secondPointer).age -= 25
	PrintStudent(*studPointer)
	studPointerNew := new(Student)
	fmt.Println("studPointer", studPointerNew)
	human := &Human{
		firstName: "John Doe",
		string:    "Addition info",
		bool:      true,
	}
	fmt.Println("anon field string ->", human.string)
	// fmt.Println(human)
}
