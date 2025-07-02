package main

import "fmt"

func main() {
	a := 10

	if a%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 is even")
	}
}
