package main

import (
	"fmt"
)

// Func expression
// assigning a function a variable

func main() {

	f := func() {
		fmt.Println("My first func expression")
	}
	f()

	g := func(x int) {
		fmt.Println("The year big brother started watching:", x)
	}
	g(1984)
}
