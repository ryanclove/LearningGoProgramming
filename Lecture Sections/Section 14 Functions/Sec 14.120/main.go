package main

import (
	"fmt"
)

// Recursion
// a function that calls itself
// just use loops, recursion is not ideal

func main() {
	fmt.Println(4 * 3 * 2 * 1)
	n := factorial(4)
	fmt.Println(n)

	// with loop
	l := loopFactorial(4)
	fmt.Println(l)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// 4 * factorial(4-1) * factorial(3-1) * factorial(2-1) * factorial(1-1)

// rewritten as a loop
func loopFactorial(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
