package main

import "fmt"

func digits(number int, dchan chan int) {
	for number != 0 {
		digit := number % 10
		dchan <- digit
		number /= 10
	}
	close(dchan)
}

func squaresGoRoutine(num int, squareChan chan int) {
	sum := 0
	dch := make(chan int)
	go digits(num, dch)
	for digit := range dch {
		sum += digit * digit
	}
	squareChan <- sum
	close(squareChan)
}
func cubesGoRoutine(num int, cubeChan chan int) {
	sum := 0
	dch := make(chan int)
	go digits(num, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeChan <- sum
	close(cubeChan)
}

func main() {
	number := 589
	squareChan := make(chan int)
	cubeChan := make(chan int)
	go squaresGoRoutine(number, squareChan)
	go cubesGoRoutine(number, cubeChan)
	squaresSum, cubesSum := <-squareChan, <-cubeChan
	fmt.Println("Total result is:", squaresSum+cubesSum)
}
