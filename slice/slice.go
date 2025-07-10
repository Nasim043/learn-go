package main

import "fmt"

func main() {
	// uninitialized slice is nil
	// var nums []int
	// fmt.Println(nums)        // []
	// fmt.Println(nums == nil) // true

	// make(type, len, cap)
	var nums1 = make([]int, 0, 2)
	fmt.Println(nums1) // []
	
	nums1 = append(nums1, 1)
	fmt.Println(nums1, cap(nums1)) // [1] 2
	
	nums1 = append(nums1, 2, 3)
	fmt.Println(nums1, cap(nums1)) // [1 2 3] 4
	nums1 = append(nums1, 2, 3, 4, 5)
	fmt.Println(nums1, cap(nums1)) // [1 2 3 2 3 4 5] 8

	// copy
	src := []int{1,2,3}
	dest := make([]int, len(src))

	// copy(dest, src)
	copy(dest, src)
	fmt.Println(dest) // [1 2 3]

}
