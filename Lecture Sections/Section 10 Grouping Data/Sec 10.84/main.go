package main

import (
	"fmt"
)

// Slice - delete from a slice

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{234, 456, 678, 987}
	x = append(x, y...)
	fmt.Println(x)

	// delete from a slice
	x = append(x[:2], x[4:]...) // needs the dots "unfurl it"
	fmt.Println(x)
	// took out elements/indexes 2 and 3

}
