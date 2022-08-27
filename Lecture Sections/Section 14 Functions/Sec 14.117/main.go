package main

import (
	"fmt"
)

// Returning a func

func main() {
	s1 := foo()
	fmt.Println(s1)

	x := bar()
	fmt.Printf("%T\n", x)
	// type is func() int
	// we are able to print by run the function
	fmt.Println(x())

	fmt.Println("-----------")

	// This also does the same thing for printing 451
	fmt.Println(bar()())
}

// return a string
func foo() string {
	return "Hello word"
}

// function (bar()) returning a function (func() int) that returns a value (451)
func bar() func() int {
	return func() int {
		return 451
	}
}
