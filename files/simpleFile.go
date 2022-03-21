package main

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read(filepath string) {
	info, err := ioutil.ReadFile(filepath)
	check(err)
	fmt.Println(string(info))
}

func write(filepath string, message string) {
	o1 := []byte(message)
	err := ioutil.WriteFile(filepath, o1, 0644)
	check(err)
}

func main() {
	read("test.txt")
	write("output.txt", "hey, dude!")
}
