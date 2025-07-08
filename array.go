package main

import "fmt"

func main() {
	var arr [2]int
	fmt.Println(arr) // [0 0]

	arr[0] = 1
	fmt.Println(arr) // [1 0]

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// declare and initialize
	arr2 := [5]int{1, 2, 3}
	fmt.Println(arr2) // [1 2 3 0 0]
}
