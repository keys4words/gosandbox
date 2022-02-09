package main

import (
	"fmt"
	"unicode/utf8"
)

type BigWord interface {
	IsBig() bool
}

type MySuperString string

func (s MySuperString) IsBig() bool {
	if utf8.RuneCountInString(string(s)) > 10 {
		return true
	}
	return false
}

func main() {
	sample := MySuperString("sldjf")
	var interfaceSample BigWord
	interfaceSample = sample
	fmt.Println("IsBig? ", interfaceSample.IsBig())

	// var interfaceBadSample BigWord
	// interfaceBadSample = "absfldfje"

}