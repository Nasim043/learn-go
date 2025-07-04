// init function : you cant call this,
// computer will call this automatically

package main

import "fmt"

var a = 10

func main() {
	fmt.Println(a)
}

func init() {
	fmt.Println("I am the first function which is executed")
	fmt.Println(a)
	a = 20
}
