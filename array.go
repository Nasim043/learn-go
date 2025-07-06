package main

import "fmt"

func main() {
	var arr [2] int
	fmt.Println(len(arr))

	for i := 0; i < len(arr); i++{
		fmt.Println(arr[i])
	}
}