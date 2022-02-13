package main

import (
	"errors"
	"fmt"
	"runtime/debug"
	"strconv"
)

func functionError(a int) (string, error) {
	if a%2 == 0 {
		return "Even", nil
	}
	return "", errors.New("This is ERROR!")
}
func PanicRecover() {
	if r := recover(); r != nil {
		fmt.Println("Panic satisfied:", r)
		debug.PrintStack()
	}
}

func panicExample(firstName *string, lastName *string) {
	defer PanicRecover()
	if firstName == nil {
		panic("runtime error: firstname can not be nill")
	}
	if lastName == nil {
		panic("runtime error: lastname can not be nil")
	}
	fmt.Println("Full name:", *firstName, *lastName)
}

func main() {
	numLiteral := "12"
	num, err := strconv.Atoi(numLiteral)
	if err != nil {
		fmt.Println("Can not convert this value to int", err)
		return
	}
	fmt.Println("Convertion done:", num)
	// ans, err := functionError(93)
	// if err != nil {
	// 	fmt.Println("not use odd value", err)
	// 	return
	// }
	// fmt.Println(ans)
	var name = "James"
	panicExample(&name, nil)
}
