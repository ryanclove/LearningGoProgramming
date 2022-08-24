package main

import (
	"fmt"
)

// Slice - append to a slice
// append works like Printf

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{234, 456, 678, 987}
	// ...T take an unlimited number of this type
	// T... put all those values appended to the end
	x = append(x, y...)
	fmt.Println(x)

}
