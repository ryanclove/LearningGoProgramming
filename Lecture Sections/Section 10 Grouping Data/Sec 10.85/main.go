package main

import (
	"fmt"
)

/* Slice - make
make is a built in function to create slices, maps, and channels and it
returns an intialized (not zeroes) value of type T
ex.
make([]int, 10, 100)
allocates an array of 100 ints then creates a slice structure with length 10 and capacity 100
at the first 10 elements of the array
*/

func main() {
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x)) // capacity

	x[0] = 42
	x[9] = 999
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x)) // capacity

	// index out of range error
	// x[10] = 3423

	// So we do this
	x = append(x, 32)
	fmt.Println(x)
	fmt.Println(len(x)) // 11
	fmt.Println(cap(x)) // 12

	x = append(x, 38)
	fmt.Println(x)
	fmt.Println(len(x)) // 12
	fmt.Println(cap(x)) // 12

	// capacity was met, and we want to append again,
	// so creates a new array that's double its size
	x = append(x, 66)
	fmt.Println(x)
	fmt.Println(len(x)) // 13
	fmt.Println(cap(x)) // 24
}
