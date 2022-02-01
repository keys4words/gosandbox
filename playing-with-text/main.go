package main

import (
	"fmt"
	"text/template"
	"os"
)

type secret struct {
	Name string
	Secret string
}

func main() {
	secret := secret{"Mike Tyson", "I'm not a ninja"}
	templatePath := "C:/Users/User/Documents/gosandbox/playing-with-text/template.txt"
	t, err := template.New("template.txt").ParseFiles(templatePath)
	if err != nil {
		fmt.Println(err)
	}
	err = t.Execute(os.Stdout, secret)
	if err != nil {
		fmt.Println(err)
	}
	
}