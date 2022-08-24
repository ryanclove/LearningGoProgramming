package main

import (
	"fmt"
)

// Slice - slicing a slice

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x[1]) // accessing by index
	// colon operator
	fmt.Println(x[1:])  // prints from position 1 to the end
	fmt.Println(x[1:3]) // prints from position 1 UP TO (not including) position 3

	for i, v := range x {
		fmt.Println(i, v)
	}

	// without range
	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}

	// without range but I'm using length function
	for i := 0; i <= len(x)-1; i++ {
		fmt.Println(i, x[i])
	}
}
