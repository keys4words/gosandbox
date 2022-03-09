package main

import (
	"errors"
	"fmt"
)

func ReturningError() error {
	err := errors.New("This my error")
	return err
}

func main() {
	err := ReturningError()
	if err != nil {
		panic(err)
	}
	fmt.Println("after func with error")
}
