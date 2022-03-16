package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("FOO", "22")
	fmt.Println("Env var FOO =", os.Getenv("FOO"))
	for _, el := range os.Environ() {
		fmt.Println(el)
	}
}
