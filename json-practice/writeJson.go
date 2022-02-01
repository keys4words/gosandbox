package main

import (
	"fmt"
	"encoding/json"
	"log"
)

type Professor struct {
	Name string `json:"name"`
	ScienceID int `json:"science_id"`
	IsWorking bool `json:"is_working"`
	University University `json:"university"`
}

type University struct {
	Name string `json:"name"`
	City string `json:"city"`
}

func main() {
	prof1 := Professor {
		Name: "Bob",
		ScienceID: 87847,
		IsWorking: true,
		University: University {
			Name: "BMSTU",
			City: "Moscow",
		},
	}

	byteArr, err := json.MarshalIndent(prof1, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(byteArr))
	os.WriteFile("prof1.json", byteArr, 0664)
}