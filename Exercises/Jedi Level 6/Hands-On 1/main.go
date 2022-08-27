package main

import (
	"fmt"
)

/*
○ create a func with the identifier foo that returns an int
○ create a func with the identifier bar that returns an int and a string
○ call both funcs
○ print out their results
*/

func main() {
	fmt.Println(foo())
	fmt.Println(bar())

	// can also do
	n := foo()
	x, s := bar()
	fmt.Println(n, x, s)
}

func foo() int {
	return 10
}

func bar() (int, string) {
	return 10, "Banana"
}
