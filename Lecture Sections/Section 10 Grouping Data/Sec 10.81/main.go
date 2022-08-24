package main

import (
	"fmt"
)

// Slice - for range

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(len(x))
	fmt.Println(x)
	fmt.Println(x[0]) // access by index position
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])
	fmt.Println("Using for range: ")

	for i, v := range x {
		fmt.Println(i, v)
	}
}
