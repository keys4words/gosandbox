package main

import "fmt"

func mutation(arr *[3]int) {
	// (*arr)[1] = 90
	// (*arr)[2] = 324093
	arr[1] = 90
	arr[2] = 343534
}

func mutationSlice(sls []int) {
	sls[1] = 900
	sls[2] = 934845
}

func main() {
	arr := [3]int{1, 2, 3}
	fmt.Println("Arr before mutation: ", arr)
	mutation(&arr)
	fmt.Println("Arr after mutation: ", arr)

	newArr := [3]int{1,2,3}
	fmt.Println("newArr before mutation: ", newArr)
	mutationSlice(newArr[:])
	fmt.Println("newArr after mutation: ", newArr)

}