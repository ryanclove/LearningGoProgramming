package main

import (
	"fmt"
)

// Slice - multi-dimensional slice

func main() {
	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)
	// we now have 2 slices

	// multi-dimensional slice
	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
