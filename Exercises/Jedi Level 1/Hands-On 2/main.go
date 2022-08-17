package main

import "fmt"

// Hands-on exercise #2
// Use var to declare variables with package level scope
// (like global variables)

var x int
var y string
var z bool

func main() {
	fmt.Println(x)
	// outputs 0
	fmt.Println(y)
	// outputs ""
	fmt.Println(z)
	// outputs false

	// aka the zero value
}
