package main

import (
	"fmt"
)

var a = 20
var b = 30

func add(x int, y int) int{
	return x + y
}

func main() {
	p := 30
	q := 40

	fmt.Println("p+q = ", add(p, q))
	fmt.Println("a+b = ", add(a, b))
	fmt.Println("a+p = ", add(a, p))
}
