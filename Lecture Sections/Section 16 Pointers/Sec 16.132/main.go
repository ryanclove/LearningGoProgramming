package main

import (
	"fmt"
)

// Pointers

// * derefences an address, giving you the value stored
// & gives us an address

func main() {
	a := 42
	fmt.Println(a)
	// Print the address in memory
	fmt.Println(&a)
	// Type of that value
	fmt.Printf("%T\n", a)
	// Type of the address in memory
	fmt.Printf("%T\n", &a) // prints *int

	// makes b a pointer to the memory address where an int is stored
	b := &a
	// prints the same address
	fmt.Println(b)
	// dereferencing an address
	fmt.Println(*b) // prints 42

	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", &a)
	// b is of type "int pointer"
	// *int -- the * is part of the type -- b is of type *int

	// derefences an address
	fmt.Println(*&a)

	// *b is pointing to the address of a, assigning the value 43
	*b = 43
	fmt.Println(a) // prints 43
}
