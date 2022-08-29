package main

import (
	"fmt"
)

// When to use Pointers

// it's all Pass by Value

// good if u have a huge chunk of data,
// you can just pass an address of a value stored

// can also be used to change something/value at a certain location

func main() {
	x := 0
	fmt.Println("Address where x is stored - before foo:\t", &x)
	fmt.Println("Value of x at that address - before foo:", x)
	foo(&x)
	fmt.Println("Address where x is stored - after foo:\t", &x)
	fmt.Println("Value of x at that address - after foo:\t", x)
}

// pass in a pointer to an int
func foo(y *int) {
	fmt.Println("Address where y is stored - before foo:\t", y)
	fmt.Println("Value of y at that address - before foo:", *y)
	*y = 43
	fmt.Println("Address where y is stored - after foo:\t", y)
	fmt.Println("Value of y at that address - after foo:\t", *y)
}
