package main

import (
	"fmt"
)

/*
Using the code from the previous example, use SLICING to create the following new slices
which are then printed:
● [1 2 3 4 5]
● [6 7 8 9 10]
● [3 4 5 6 7]
● [2 3 4 5 6]
*/

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// remember, it's index value for x[:]

	firstFive := x[:5]
	fmt.Println(firstFive)

	lastFive := x[5:]
	fmt.Println(lastFive)

	threeToSeven := x[2:7]
	fmt.Println(threeToSeven)

	twoToSix := x[1:6]
	fmt.Println(twoToSix)

}
