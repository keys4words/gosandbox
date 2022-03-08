package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int = 44
	fmt.Println("a is", unsafe.Sizeof(a), "bytes")

}
