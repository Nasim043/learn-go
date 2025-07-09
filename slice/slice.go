package main

import "fmt"

func main() {
	// uninitialized slice is nil
	var nums []int
	fmt.Println(nums)        // []
	fmt.Println(nums == nil) // true
}
