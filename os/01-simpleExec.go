package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("firefox")
	err := cmd.Run()
	if err != nil {
		log.Fatal()
	}
}
