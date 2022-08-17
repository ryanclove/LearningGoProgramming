package main

import "fmt"

// Bool type

func main() {
	a := 7
	b := 42
	fmt.Println(a == b)
	// Operator checks for same address comparison, Prints false
	fmt.Println(a <= b)
	// Prints true
	fmt.Println(a >= b)
	// Prints false
	fmt.Println(a != b)
	// Prints true
}

// = assignment operator
// == quality comparison
