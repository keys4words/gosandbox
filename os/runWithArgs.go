package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func RunWithArgs() {
	cmd := exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("my test string")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("translated phrase: %q\n", out.String())
}

func RunWithManyArgs() {
	prog := "echo"
	arg1 := "there"
	arg2 := "must be"
	arg3 := "your ads"
	cmd := exec.Command(prog, arg1, arg2, arg3)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Print(string(stdout))
}

func main() {
	// cmd := exec.Command("firefox")
	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatal()
	// }
	// RunWithArgs()
	RunWithManyArgs()
}
