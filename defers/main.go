package main

import "fmt"

func CheckDBCloseConnection(a int) {
	fmt.Println("Check started... Value numIP in defer:", a)
	fmt.Println("Check done! Connection refused!")
}
func OuterFunc() int {
	defer fmt.Println("I'm deferred print function!")
	fmt.Println("OuterFunc started")
	var result = 100
	fmt.Println("OuterFunc finished. Ready to return value!")
	return result
}

func main() {
	var numIP = 10
	p := &numIP
	// defer CheckDBCloseConnection(numIP)
	*p++
	fmt.Println("Value numIP in main():", numIP)
	fmt.Println("Main function started")
	fmt.Println("Main function ended!")
	fmt.Println("Value form OuterFunc on main side is:", OuterFunc())
}
