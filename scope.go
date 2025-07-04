package main 

import (
	"fmt"

	"example.com/mathlib"
)

var a = 20
var b = 30

func add(x int, y int) int {
	return x + y
}

func main() {
	p := 30
	q := 40

	fmt.Println("p+q = ", add(p, q))
	fmt.Println("a+b = ", add(a, b))
	fmt.Println("a+p = ", add(a, p))

	// package level scope
	// go mod init example.com
	fmt.Println(mathlib.Add(p, q))
}
