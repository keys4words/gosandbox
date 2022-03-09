package main

import (
	"fmt"
	"strings"
)

func main() {
	var password string
	for {
		fmt.Print("Insert password: ")
		fmt.Scan(&password)
		if strings.Contains(password, "123") {
			fmt.Println("weak password")
			continue
		} else {
			fmt.Println("Password accepted")
			return
		}

	}
}
