package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while loop
	j := 0
	for j < 2 {
		fmt.Println(j)
		j++
	}

	// infinite loop
	for {
		fmt.Println("loop")
		break
	}
	// range
	for k := range 3 {
		fmt.Println("range", k)
	}

	// nested loop
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(i, j)
		}
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
