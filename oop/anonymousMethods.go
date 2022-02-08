package main

import "fmt"

type University struct {
	city string
	name string
}

func (u *University) FullInfoUniversity() {
	fmt.Printf("Univer name: %s and city: %s\n", u.name, u.city)
}

type Professor struct {
	fullName string
	age int
	University
}

func main() {
	p := Professor{
		fullName: "Boris Bobroff",
		age: 150,
		University: University{
			city: "Moscow",
			name: "BMSTU",
		},
	}
	p.FullInfoUniversity()
}