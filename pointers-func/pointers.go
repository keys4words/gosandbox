package main

import "fmt"

func changeParam(val *int){
	*val += 100
}

func returnPointer() *int {
	var numeric int = 34
	return &numeric
}

func main() {
	// var variable int = 30
	// var point *int = &variable
	// fmt.Printf("type: %T and value: %v\n", point, point)

	// var zeroPointer *int
	// fmt.Printf("type: %T and value: %v\n", zeroPointer, zeroPointer)
	// if zeroPointer == nil {
	// 	zeroPointer = &variable
	// }
	// fmt.Printf("After: type: %T and value: %v\n", zeroPointer, zeroPointer)
	// var numericValue int = 32
	// var pointerToNumeric *int = &numericValue
	// fmt.Printf("Value in numericValue is %v\nAddress is %v\n", *pointerToNumeric, pointerToNumeric)

	// var zeroVar int
	// var zeroPoint *int = &zeroVar
	// fmt.Printf("Value in *zeroPointer %v\nAddress is %v\n", *zeroPoint, zeroPoint)

	// zeroPoint2 := new(int)
	// fmt.Printf("Value in *zeroPoint2 %v\nAddress is %v\n", *zeroPoint2, zeroPoint2)

	zeroPointerToInt := new(int)
	fmt.Printf("Address is %v and Value in zeroPointerToInt is %v\n", zeroPointerToInt, *zeroPointerToInt)
	*zeroPointerToInt += 430
	fmt.Printf("Address is %v and Value in zeroPointerToInt is %v\n", zeroPointerToInt, *zeroPointerToInt)
	b := 34
	a := &b
	*a++
	fmt.Println(b)

	sample := 1
	samplePointer := &sample
	fmt.Println("Origin value of sample: ", sample)
	changeParam(samplePointer)
	fmt.Println("After changing: ", sample)

	ptr1 := returnPointer()
	ptr2 := returnPointer()
	fmt.Printf("Ptr1: %T and address %v and value %v\n", ptr1, ptr1, *ptr1)
	fmt.Printf("Ptr2: %T and address %v and value %v\n", ptr2, ptr2, *ptr2)

}