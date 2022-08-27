package main

import (
	"fmt"
)

// Closure
// we want to close the scope of a variable to contain it

// package level scope
var x int

func main() {
	fmt.Println(x)
	x++
	fmt.Println(x)
	foo()

	// code block level scope
	{
		y := 42
		fmt.Println(y)
	}
	// printing y outside the block would result in error
	// fmt.Println(y)

	// func main() level scope
	var z string = "z"
	fmt.Println(z)

	// return incrementer() value which returns func() int in its scope
	b := incrementor()
	c := incrementor()
	fmt.Println(b(), c())
	fmt.Println(b(), c())
	fmt.Println(b(), c())
}

func foo() {
	fmt.Println("Entering foo:", x)
	x++
	fmt.Println("Leaving foo:", x)

	// z would not pass into foo, cause an error b/c it's scope is only in main()
	// z = append(z, z)
}

func incrementor() func() int {
	var a int
	return func() int {
		a++
		return a
	}
}
