package main

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



func main() {

}