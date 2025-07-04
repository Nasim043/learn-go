package main

import "fmt"

var a = 10

func main() {
	age := 30

	if age > 18 {
		//  a variable temporarily "hides" or shadows
		// the outer a variable in this block
		a := 47
		fmt.Println(a) // 47
	}

	fmt.Println(a) // 10
}

// When two different memory locations are created with same name,
// and one shadows another, then it's variable shadowing.

// Variable shadowing happens when a local variable
// (usually inside a block or function)
// has the same name as a variable in an outer scope,
// and temporarily "hides" or shadows the outer variable in that block.

// The outer variable is still there, but you can’t access it
// from within the block where the new variable shadows it.
