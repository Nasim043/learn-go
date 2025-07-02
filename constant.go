package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s) // constant

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d) // 6e+11

	// A numeric constant has no type until it’s given one
	fmt.Println(int64(d)) // 600000000000

	fmt.Println(math.Sin(d)) // 0.7591864109375384
}
