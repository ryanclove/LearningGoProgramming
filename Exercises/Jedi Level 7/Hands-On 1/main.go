package main

import (
	"fmt"
)

/*
● Create a value and assign it to a variable.
● Print the address of that value
*/

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)
}
