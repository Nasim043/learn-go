package main

import "fmt"

func main() {
	// initializaion
	arr := [3]int{1, 2, 3}
	fmt.Println(arr) // [1 2 3]

	// length
	fmt.Println(len(arr)) // 3

	// loop
	for i, v := range arr {
		fmt.Println(i, v)
	}

	// modify
	modifyArray(arr)
	fmt.Println(arr)
}

// doesn't change original due to pass-by-value
func modifyArray(arr [3]int) {
	arr[1] = 20      // This modification is local to the function
	fmt.Println(arr) // [1 20 3]
}
