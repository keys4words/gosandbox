package main

import "fmt"

type University struct {
	age int
	yearBased int
	infoShort string
	infoLong string
}

type Student struct {
	firstName string
	lastName string
	university University
}

type Professor struct {
	firstName string
	lastName string
	age int
	greatWork string
	University
}

func main() {
	stud := Student {
		firstName : "Fedya",
		lastName: "Petrov",
		university: University {
			yearBased: 1991,
			infoShort: "cool University",
			infoLong: "very cool university",
		},
	}

	fmt.Println(stud.firstName, stud.university.yearBased)

	prof := Professor{
		firstName: "Petr",
		lastName: "Schtamm",
		age: 125,
		greatWork: "Ultimate Golang",
		University: University{
			yearBased: 1765,
			infoShort: "long time ago",
			age: 2021-1765,
		},
	}
	fmt.Println(prof.infoShort, prof.age, prof.University.age)
	profLeft := Professor{
		firstName: "Nilse Bor",
	}
	profRight := Professor{}
	fmt.Println(profLeft == profRight)
}