package main

import (
	"fmt"
)


func main() {
	mapper := make(map[string]int)
	mapper["Alice"] = 435
	mapper["Joe"] = 555
	fmt.Println("map: ", mapper)

	newMapper := map[string]int{
		"Alice": 1000,
		"Cuper": 2000,
	}
	newMapper["Cuper"] = 3000
	fmt.Println(newMapper)
	// var mapZeroValue map[string]int
	// mapZeroValue["Alice"] = 999
	testPerson := "Alice"
	fmt.Println("Salary = ", newMapper[testPerson])
	unexistPerson := "john"
	fmt.Println("unexistPerson: ", newMapper[unexistPerson])
	if value, ok := newMapper[unexistPerson]; ok {
		fmt.Println("unexistPerson: ", value)
	} else {
		fmt.Println("there is no this person")
	}
	newMapper["James"] = 9997
	delete(newMapper, "Jares")	//not check if "Jares" is exists

	fmt.Println("================")
	for k, v := range newMapper {
		fmt.Printf("%s and value %d\n", k, v)
	}
	words := map[string]string {
		"one": "Один",
		"two" : "Два",
	}
	newWords := words
	newWords["Three"] = "Три"
	delete(newWords, "one")
	fmt.Println("words: ", words)
	fmt.Println("newWords: ", newWords)

	if [3]int{1,2,3} == [3]int{1,2,3}{
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
	// if []int{1,2,3} == []int{1,2,3}{
	// 	fmt.Println("equal")
	// }
		
}