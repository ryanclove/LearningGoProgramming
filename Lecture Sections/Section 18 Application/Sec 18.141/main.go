package main

import (
	"fmt"
	"sort"
)

// Sort

func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}
	fmt.Println()

	// does not need a declaration or expression, like xi = sort.Ints(xi)
	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println("Sorted using sort package:")
	fmt.Println(xi)

	fmt.Println("---------------------")

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println("Sorted using sort package:")
	fmt.Println(xs)
	fmt.Println()

}
