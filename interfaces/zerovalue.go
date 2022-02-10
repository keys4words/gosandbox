package main

import "fmt"

type Rider interface {
	Ride()
	Gas()
	Stop()
}


func main() {
	var r Rider
	if r == nil {
		fmt.Printf("r is nill. Value %v and type %T\n", r, r)
	}
	// r.Ride()
}