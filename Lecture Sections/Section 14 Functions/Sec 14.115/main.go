package main

import (
	"fmt"
)

// Anonymous functions

func main() {
	foo()

	// anonymous function does not have a name
	func() {
		fmt.Println("Anonymous func ran")
	}() //parens at the end too

	func(x int) {
		fmt.Println("The meaning of life:", x)
	}(42) // pass in that argument
}

func foo() {
	fmt.Println("foo ran")
}
