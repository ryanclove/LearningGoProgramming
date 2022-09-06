package main

import (
	"fmt"
)

/*
Using a COMPOSITE LITERAL:
	● create an ARRAY which holds 5 VALUES of TYPE int
	● assign VALUES to each index position.
Range over the array and print the values out.
Using format printing
	● print out the TYPE of the array
*/

func main() {
	var x [5]int
	// instructor just declared all 5 values
	// x := [5]int{40, 41, 42, 43, 44}

	// I decided to use a for loop for fun
	for i := 0; i < len(x); i++ {
		x[i] = i
	}

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("Array is of type %T\n", x)
}
