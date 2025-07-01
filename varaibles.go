package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a) // initial

	var b, c int = 1, 2
	fmt.Println(b, c) // 1 2

	var d = true
	fmt.Println(d) // true

	// Variables declared without a corresponding initialization are zero-valued.
	var e int
	fmt.Println(e) // 0

	// The := syntax is shorthand for declaring and initializing a variable.
	// This syntax is only available in functions.
	// var f string = "apple"
	f := "apple"
	fmt.Println(f) // apple
}
