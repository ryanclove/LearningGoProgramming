package main

import (
	"fmt"
)

// Callback
// passing a func as an argument

// called functional programming, not reccomended

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(ii...)
	fmt.Println("Sum of all numbers:", s)

	s2 := even(sum, ii...)
	fmt.Println("Sum of even numbers:", s2)

	s3 := odd(sum, ii...)
	fmt.Println("Sum of odd numbers:", s3)
}

func sum(xi ...int) int {
	// fmt.Printf("%T\n", xi)
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

// callback  // passes in a variadic parameter func int and returns an int
// which of vi are even?
func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

// callback returning odd numbers
func odd(f func(xi ...int) int, vali ...int) int {
	var si []int
	for _, v := range vali {
		if v%2 != 0 {
			si = append(si, v)
		}
	}
	return f(si...)
}
