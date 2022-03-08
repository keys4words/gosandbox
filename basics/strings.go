package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	word := "Тестовая строка"
	fmt.Printf("String %s, len:%d, Rune len:%d\n", word, len(word), utf8.RuneCountInString(word))
	fmt.Printf("Bytes: ")
	for i := 0; i < len(word); i++ {
		fmt.Printf("%x ", word[i])
	}
	// fmt.Println()
	// for i := 0; i<len(word); i++ {
	// 	fmt.Printf("%c ", word[i])
	// }
	// fmt.Println()
	// runeSlice := []rune(word)
	// for i := 0; i<len(runeSlice); i++ {
	// 	fmt.Printf("%c ", runeSlice[i])
	// }
	// fmt.Println()
	// myByteSlice := []byte{0x40, 0x41, 0x42}
	// myStr := string(myByteSlice)
	// fmt.Println(myStr)
	// myDecimalBytesSlice := []byte{101, 102, 103}
	// fmt.Println(string(myDecimalBytesSlice))
	word1, word2 := "Вася", "Петя"
	if word1 == word2 {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}
	firstName := "Boris"
	lastName := "Jonson"
	fullName := fmt.Sprintf("%s ### %s", firstName, lastName)
	fmt.Println("FullName: ", fullName)
	fullNameSlice := []rune(fullName)
	fullNameSlice[0] = 'L'
	fullName = string(fullNameSlice)
	fmt.Println("String mutation: ", fullName)
}
