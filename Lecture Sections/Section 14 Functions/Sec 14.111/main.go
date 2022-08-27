package main

import (
	"fmt"
)

// Unfurling a slice
// Unfurling is not the official term for Go

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9} // a slice of int
	x := sum(xi...)                     // allows to pass to sum(...int), which Todd calls unfurling
	fmt.Println("The total is", x)

	fmt.Println("----sum function with multiple args----")

	s := sumWithStrings("James", 4, 5, 6)
	fmt.Println("The total is", s)

	fmt.Println("----sum function with no (slice of) ints passed----")

	ni := sumNoInts("James")
	fmt.Println("The total is", ni)
}

// accepts ints, not a slice of ints
// so in x above, adding the dots allows the slice values to pass
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position,", i, " we now add", v, "Which equals:", sum)
	}
	return sum
}

// varuiadic a...int must be last argument here
func sumWithStrings(str string, a ...int) int {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	sum := 0
	for i, v := range a {
		sum += v
		fmt.Println("for item in index position,", i, " we now add", v, "Which equals:", sum)
	}
	return sum
}

// ni declaration only has a string arg, so int is zero value 0
func sumNoInts(james string, r ...int) int {
	fmt.Println(r)
	fmt.Printf("%T\n", r)
	fmt.Println(len(r))
	fmt.Println(cap(r))

	sum := 0
	for i, v := range r {
		sum += v
		fmt.Println("for item in index position,", i, " we now add", v, "Which equals:", sum)
	}
	return sum
}
