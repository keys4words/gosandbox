package main

import (
	"encoding/json"
	"fmt"
)

type Dimensions struct {
	Height int
	Width  int
}

type Bird struct {
	Species     string
	Description string
	Dimensions  Dimensions
}

func ParseJson() {
	birdJson := `{"species":"pigeon","description":"likes to perch on rocks", "dimensions":{"height":24,"width":10}}`
	var bird Bird
	err := json.Unmarshal([]byte(birdJson), &bird) //Unmarshall json birdJson to bird structure

	if err != nil {
		panic(err)
	}

	fmt.Println(bird)
}

func CreateJson() {
	bird := Bird{
		Species:     "Eagle",
		Description: "Cool eagle",
		Dimensions: Dimensions{
			Height: 100,
			Width:  50,
		},
	}
	data, _ := json.MarshalIndent(bird, "", "  ") //Marshalling with indent
	fmt.Println(string(data))
}

func ParseJsonArrays() {
	arraysJson := `["one","two","three"]`

	var nums []string

	json.Unmarshal([]byte(arraysJson), &nums) //JSON parse to slice

	fmt.Println(nums)
	fmt.Println(nums[0])
}

func main() {
	// ParseJson()
	// CreateJson()
	ParseJsonArrays()
}
