/*
	1. parameter vs argument
	2. first order function
		i. standard function or named function
		ii. anonymous function
		iii. IIFE
		iv. function expression
	3. higher order funtion
		one of the following
		i. function as a parameter
		ii. function as a return
		iii. both
	4. callback function
	the function which is passed as an argument to another function
*/

package main

import "fmt"

func sum(a int, b int) { // parameter => a,b
	c := a + b
	fmt.Println(c)
}

func processOperation(a int, b int, operation func(p int, q int)) {
	operation(a, b) // 7
}

func main() {
	sum(2, 5) // argument => 2,5

	processOperation(2, 5, sum)
}
