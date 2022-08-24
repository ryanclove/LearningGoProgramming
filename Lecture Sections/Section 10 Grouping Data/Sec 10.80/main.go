package main

import (
	"fmt"
)

// Slice - composite literal

func main() {
	//x := type{values} // composite literal
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
}

// Slice allows you to group together VALUES of the same TYPE
