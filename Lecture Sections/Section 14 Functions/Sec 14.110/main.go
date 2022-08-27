package main

import (
	"fmt"
)

// Variadic parameter

func main() {
	x := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The total is", x)
}

// func (r receiver) indentifier(parameter(s)) (return(s)) {code}

// ...int means unlimited  number of ints
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x) // type is a slice of []int

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position,", i, " we now add", v, "Which equals:", sum)
	}
	return sum
}
