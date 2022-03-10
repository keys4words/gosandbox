package main

import "fmt"

func main() {

	// originalArr := [...]int{30, 40, 50, 60, 70, 80}
	// firstSlice := originalArr[1:4]
	// for i, _ := range firstSlice {
	// 	firstSlice[i]++
	// }
	// fmt.Println("OriginalArr: ", originalArr)
	// fSlice := originalArr[:]
	// fSlice = append(fSlice, 59)
	// // capacity - это значение, показывающее сколько элементов в принципе можно добавить в срез без выделения дополнительной памяти
	// fmt.Println("length: ", len(fSlice), " capacity: ", cap(fSlice))
	// fmt.Println(fSlice)
	// numArr := [10]int{1, 2}
	// numSlice := numArr[:]
	// numSlice = append(numSlice, 3)
	// numSlice[0] = 10
	// fmt.Println(numArr)
	// fmt.Println(numSlice)

	s1 := make([]int, 10, 16)
	fmt.Println(s1)
	myWords := []string{"one", "two", "three"}
	fmt.Println("My words: ", myWords)
	// anotherSlice := []string{"four", "five", "six"}
	// myWords = append(myWords, anotherSlice...)
	// fmt.Println("My words: ", myWords)

}
