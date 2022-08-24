package main

import (
	"fmt"
)

// Array
// We don't really use arrays in GO, they're primarily
// building blocks for slices - use slices

func main() {
	var x [5]int
	var y [6]int
	fmt.Println(x)
	x[4] = 42
	fmt.Println(x)
	fmt.Println(y)
	// length is part of an array's type
	fmt.Println(len(x))
	fmt.Println(len(y))
}
