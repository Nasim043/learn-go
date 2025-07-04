package main

import "fmt"

func main() {
	// anonymous function
	// Immediately Invoked Function Expression
	// IIFE
	func(a int, b int) {
		c := a + b
		fmt.Println(c)
	}(10, 5)
}

func init() {
	fmt.Println("I am the first function which is executed")
}
